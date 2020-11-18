package metrics

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
	"sync"
)

var (
	DefEtcdName = "/metrics/registry/"
)

// 指标注册
type RegisterMetrics struct {
	s        *ServiceRegister
	ri       *RegisterInfo
	listener net.Listener
}

type RegisterInfo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// 创建 注册信息
func NewRegisterMetrics(name string) *RegisterMetrics {
	return &RegisterMetrics{
		ri: &RegisterInfo{
			Name: name,
		}}
}

// 关闭注册信息
func (r *RegisterMetrics) Close() {
	if r == nil {
		return
	}
	if r.s != nil {
		_ = r.s.Close()
	}
	if r.listener != nil {
		_ = r.listener.Close()
	}
	log.Printf("%v for etcd closed ... ", r.ri.Name)
}

// 开始注册信息
//
// etcd 的服务地址
//
// endpoints ："localhost:2379","192.168.7.51:2379"
func (r *RegisterMetrics) Register(endpoints ...string) error {
	if len(endpoints) == 0 {
		log.Printf("etcd endpoints is nil ,the metrics will not be available ... ")
		endpoints = append(endpoints, "localhost:2379")
	}
	// 序列化
	marshal, _ := json.Marshal(r.ri)

	if r.s == nil {
		ser, err := NewServiceRegister(endpoints, DefEtcdName+r.ri.Name+"", string(marshal), 5)
		if err != nil {
			log.Printf("new etcd failed ... ,err = %v", err)
			return err
		}
		r.s = ser
		if ser == nil && r.listener != nil {
			r.listener.Close()
			return err
		}
		//监听续租相应chan
		go func() {
			r.s.ListenLeaseRespChan()
			log.Printf("%v close renewal ... ", r.ri.Name)
		}()
	}
	return nil
}

var o sync.Once

// 开启指标并注册到 etcd 中。
//
// metricsAddr 默认给空 则随机端口。
func (r *RegisterMetrics) StartMetrics(etcdAddr, metricsAddr string) error {
	var err error
	o.Do(func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		r.listener, err = net.Listen("tcp", metricsAddr)
		if err != nil {
			log.Printf("err : %v", err)
			return
		}
		go func() {
			err = http.Serve(r.listener, mux)
			if err != nil {
				log.Printf("err : %v", err)
			}
		}()
		log.Printf("name : %v , addr : %v  ready to open metrics ", r.ri.Name, r.listener.Addr().String())

		// 提取真实 ip 地址
		var ip, host, port string
		// ipv6 address in format [host]:port or ipv4 host:port
		host, port, err = net.SplitHostPort(r.listener.Addr().String())
		if err != nil {
			log.Printf("err : %v", err)
		}
		ip, err = Extract(host)
		if err != nil {
			log.Printf("err : %v", err)
		}
		addr := HostPort(ip, port)

		// 设置监控地址
		r.ri.Address = addr

		// 开启
		if err = r.Register(etcdAddr); err != nil {
			log.Printf("Register [etcd] failed %v ", err)
		}
	})
	return err
}
