package log

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	log := zap.NewExample()
	log.Debug("this is debug message")
	log.Info("this is info message")
	log.Info("this is info message with fileds",
		zap.Int("age", 24), zap.String("agender", "man"))
	log.Warn("this is warn message")
	log.Error("this is error message")
}

func TestDevelopment(t *testing.T) {
	log, _ := zap.NewDevelopment()
	log.Debug("this is debug message")
	log.Info("this is info message")
	log.Info("this is info message with fileds",
		zap.Int("age", 24), zap.String("agender", "man"))
	log.Warn("this is warn message")
	log.Error("this is error message")
}

func TestProduction(t *testing.T) {
	log, _ := zap.NewProduction()
	log.Debug("this is debug message")
	log.Info("this is info message")
	log.Info("this is info message with fileds",
		zap.Int("age", 24), zap.String("agender", "man"))
	log.Warn("this is warn message")
	log.Error("this is error message")
}

func TestNormal(t *testing.T) {

	L.Debug("init...")

	time.Sleep(time.Millisecond)
	L.Info("init...")
	time.Sleep(time.Millisecond)
	L.Info("init...")
	time.Sleep(time.Millisecond)
	L.Info("init...")
	time.Sleep(time.Second)
	L.Info("init...")
	L.Info("init...")
	L.Info("init...")
	L.Info("init...")
	L.Error("init....")
	L.DPanic("init....")

}

func TestSwitchLogFile(t *testing.T) {

	for i := 0; i < 100000; i++ {
		L.Info("init...", zap.Int("idx", i))
	}

	time.Sleep(time.Second)
}
