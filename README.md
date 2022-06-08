# golearning
### 分享一些自己所遇到或者所见到的一些代码

- 项目架构很早，以前使用C的rsa来处理，我使用了go标准库的rsa用来压测对比了一下。 [Go和 C的 Rsa对比](https://github.com/ywanbing/golearning/tree/master/rsaCAndGoBaench "Go和 C的 Rsa对比")
- 如果50个协程去访问全局变量并且进行+1的操作，最终这个全局变量会不会是50呢? [查看解决方案](https://github.com/ywanbing/golearning/blob/master/usualQuestion/atomic_test.go "查看解决方案")
- go使用redis实现分布式锁，快速上手 [go-redis 分布式锁](https://github.com/ywanbing/golearning/tree/master/go-redis-distributed-lock "go使用redis实现分布式锁")
- zap 日志框架添加日志自动切换功能 [zap-log](https://github.com/ywanbing/golearning/tree/master/zap-log "zap 自动切换")
- 实现 AES-CBC 加密 [AES_CBC](https://github.com/ywanbing/golearning/tree/master/aesCrypto "AES-CBC 加密")
- 实现golang项目的社区标准目录建立 [project-layout](https://github.com/ywanbing/golearning/tree/master/script "project-layout")
- 实现 etcd + confd + prometheus 自动发现 [etcd_confd_prometheus](https://github.com/ywanbing/golearning/tree/master/etcd_confd_prometheus "etcd_confd_prometheus")
- 实现golang脚本调用，使用 c 语言的接口，为什么使用的原因，懂的都懂 [go_c_script](https://github.com/ywanbing/golearning/tree/master/goloang_c_script "goloang_c_script")
- 实现 tcp 文件传输工具，在局域网可达超快速度，支持超大文件传输，且只使用很小的内存 [file_transport](https://github.com/ywanbing/golearning/tree/master/file_transport "file transport") 
- 修改 goimports 只区分官方和非官方包的分组 [goimports](https://github.com/ywanbing/golearning/tree/master/goimports "goimports") 
- 通过 go 编写一个程序来实现 cpu 使用率在一定的时间内维持 50%。 [cpu50](https://github.com/ywanbing/golearning/tree/master/cpu50 "cpu50")
- 学习`chan`的源码，通过`chan`原理流程图，事半功倍。[chan_flowChart](https://github.com/ywanbing/golearning/tree/master/chan_flow_chart "chan_flow_chart")
- `golangci-lint`的中文规则介绍，通过`golangci-lint.yml`配置的默认规则。[golangci-lint](https://github.com/ywanbing/golearning/tree/master/golangci-lint "golangci-lint")
- 实现互斥锁的额外功能[lock](https://github.com/ywanbing/golearning/tree/master/lock "lock")

### 问题
希望有相关的问题可以提交issues中一起学习和实现
