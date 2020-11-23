# golearning
### 分享一些自己所遇到或者所见到的一些代码

- 项目架构很早，以前使用C的rsa来处理，我使用了go标准库的rsa用来压测对比了一下。 [Go和 C的 Rsa对比](https://github.com/ywanbing/golearning/tree/master/rsaCAndGoBaench "Go和 C的 Rsa对比")
- 如果50个协程去访问全局变量并且进行+1的操作，最终这个全局变量会不会是50呢? [查看解决方案](https://github.com/ywanbing/golearning/blob/master/usualQuestion/atomic_test.go "查看解决方案")
- go使用redis实现分布式锁，快速上手 [go-redis 分布式锁](https://github.com/ywanbing/golearning/tree/master/go-redis-distributed-lock "go使用redis实现分布式锁")
- zap 日志框架添加日志自动切换功能 [zap-log](https://github.com/ywanbing/golearning/tree/master/zap-log "zap 自动切换")
- 实现 AES-CBC 加密 [AES_CBC](https://github.com/ywanbing/golearning/tree/master/aesCrypto "AES-CBC 加密")
- 实现golang项目的社区标准目录建立 [project-layout](https://github.com/ywanbing/golearning/tree/master/script "project-layout")
- 实现 etcd + confd + prometheus 自动发现 [etcd_confd_prometheus](https://github.com/ywanbing/golearning/tree/master/etcd_confd_prometheus "etcd_confd_prometheus")
### 问题
希望有相关的问题可以提交issues中一起学习和实现
