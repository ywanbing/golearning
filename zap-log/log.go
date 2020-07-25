package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

const (
	PATH = "./logs/"
	NAME = "default.log"
)

var L *zap.Logger

func init() {
	ws := initWriter()
	encoder := initEncoder()

	atomicLevel := zap.NewAtomicLevel()

	//默认debug
	atomicLevel.UnmarshalText([]byte("debug"))

	//atomicLevel.SetLevel()

	core := zapcore.NewCore(
		encoder,
		ws,
		atomicLevel,
	)
	zapOptions := make([]zap.Option, 0)
	zapOptions = append(zapOptions, zap.AddStacktrace(zap.DPanicLevel), zap.AddCaller(), zap.AddCallerSkip(0))
	zapOptions = append(zapOptions, zap.Fields(zap.String("svr", "test")))
	zapOptions = append(zapOptions, zap.Development())

	L = zap.New(core, zapOptions...)
	L.Info("DefaultLogger init success")
}

func initWriter() zapcore.WriteSyncer {
	lLog := &lumberjack.Logger{
		Filename:   PATH + NAME,
		MaxSize:    1,
		MaxAge:     0,
		MaxBackups: 10,
		LocalTime:  false,
		Compress:   false,
	}
	return zapcore.AddSync(lLog)
}

func initEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.LevelKey = "lv"

	encoderConfig.EncodeTime = DiyTimeEncoder // 修改时间编码器

	// 在日志文件中使用小写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

	// NewJSONEncoder 日志格式 好分析
	return zapcore.NewJSONEncoder(encoderConfig)
}

func DiyTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
