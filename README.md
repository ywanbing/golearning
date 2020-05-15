# golearning
### 分享一些自己所遇到或者所见到的一些代码

- 项目架构很早，以前使用C的rsa来处理，我使用了go标准库的rsa用来压测对比了一下。 [Go和 C的 Rsa对比](https://github.com/ywanbing/golearning/tree/master/rsaCAndGoBaench "Go和 C的 Rsa对比")
- 如果50个协程去访问全局变量并且进行+1的操作，最终这个全局变量会不会是50呢? [查看解决方案](https://github.com/ywanbing/golearning/blob/master/usualQuestion/atomic_test.go "查看解决方案")
- go使用redis实现分布式锁，快速上手 [go-redis 分布式锁](https://github.com/ywanbing/golearning/tree/master/go-redis-distributed-lock "go使用redis实现分布式锁")