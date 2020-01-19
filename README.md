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

注意要点：

==服务注册IP地址一定不要与etcd的IP地址重合==

#### 1.3 使用文档

设置版本： ETCDCTL_API=3

使用命令:

```shell
docker exec -t etcd1 etcdctl 
```

查看Etcd版本:

```shell
docker exec -t etcd1 etcdctl version
```

查看成员:

```shell
docker exec -t etcd1 etcdctl member list
```

查看集群状态:

```shell
docker exec -t etcd1 etcdctl cluster-health
```

支持远端操作：

```shell
docker exec -t etcd1 etcdctl --endpoints=127.0.0.1:2379
```

设置 Put:

```shell
docker exec -t etcd1 etcdctl put /root/test/ "hello word"
```

获取 Get:

```shell
docker exec -t etcd1 etcdctl get /root/test
```

更新 Update:

```shell
docker exec -t etcd1 etcdctl update /root/test "shu min"
```

删除 Del:

```shell
docker exec -t etcd1 etcdctl del /root/test
```

观察者[监听]:

```shell
# 端口1
docker exec -t etcd1 etcdctl watch /root/test/key -f 

# 端口2
docker exec -t etcd1 etcdctl update /root/test/key "sobot" 
```

==删除监听的值后，监听仍然在监听该key值==

申请租约[过期]:

```she 
docker exec -t etcd1 etcdctl lease grant 100
```

授权租约:

```shell
docker exec -t etcd1 etcdctl put --lease=12f76fbcfd255003 /key/1 1
```

租约续约:

```shell
docker exec -t etcd1 etcdctl lease keep-alive  # 续约
```

