#### 1.1 go-micro

```shell
# 创建服务
micro new --type srv ibook/server/book 
```

#### 1.2 搭建etcd集群

```shell
$ docker pull quay.io/coreos/etcd
$ docker-compose up
$ docker ps 
# 验证集群
$ curl -L http://127.0.0.1:32787/v2/members
$ curl -L http://127.0.0.1:32789/v2/members
$ curl -L http://127.0.0.1:32791/v2/members
# 也可以使用etcdctl：
$ docker exec -t etcd1 etcdctl member list
```



