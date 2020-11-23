## promethoues 的自动发现
使用 confd 拉取 etcd 的注册信息，自动更新 promethoues 的配置。

1. 安装 etcd  
    ```shell script
    https://github.com/etcd-io/etcd/releases
    ```
2. 安装confd 
    ```shell script
    https://github.com/kelseyhightower/confd/releases
    ```
3. 安装 promethoues 
    ```shell script
    https://github.com/prometheus/prometheus/releases
    ```
上面的都是自己根据情况选择版本下载，并先把 etcd 运行起来。

根据 etcd 的文件夹内的方式注册信息到etcd中。

### 启动 confd 
设置 confd 读取的配置和模版的文件夹，并监听 etcd 的变化。
```shell script
 confd -confdir ./confd/ -config-file ./confd/conf.d/prometheus.json.toml -backend etcdv3  -watch -node http://127.0.0.1:2379  &
```
发现下面错误是和正常的因为还没有启动 promethoues 。
```shell script
curl: (7) Failed to connect to 127.0.0.1 port 9090: Connection refused
```
### 启动 promethoues 

通过 prometheus.yml 覆盖 /etc/prometheus/prometheus.yml 的文件，没有就复制过去。   
创建一个保存数据的文件夹 
```shell script
mkdir /prometheus
```
在 prometheus 的 bin 下面启动: 
```shell script
prometheus --web.enable-lifecycle \
    --config.file=/etc/prometheus/prometheus.yml \
    --storage.tsdb.path=/prometheus
```

这个时候就完成了，你可以增加或者减少服务来查看 promethoues 的动态变化。
