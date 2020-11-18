package metrics

import "testing"

/*
// 查看注册信息
etcdctl get --prefix /metrics/registry/

/metrics/registry/test
{"name":"test","address":"127.0.0.1:34119"}
/metrics/registry/test1
{"name":"test1","address":"127.0.0.1:37217"}
/metrics/registry/test2
{"name":"test2","address":"127.0.0.1:40917"}
*/

// 注册单个节点
func TestEtcd(t *testing.T) {
	r := NewRegisterMetrics("test")
	// addr 不给参数，就是随机端口。
	_ = r.StartMetrics("127.0.0.1:2379", "")

	select {}
}

// 注册多个节点信息
func TestEtcdMultiple(t *testing.T) {
	{
		r := NewRegisterMetrics("test1")
		// addr 不给参数，就是随机端口。
		_ = r.StartMetrics("127.0.0.1:2379", "")
	}
	{
		r := NewRegisterMetrics("test2")
		// addr 不给参数，就是随机端口。
		_ = r.StartMetrics("127.0.0.1:2379", "")
	}
	select {}
}
