package main

import (
	//"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"ibook/server/book/handler"
	book "ibook/server/book/proto/book"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {

	/*
		1. 初始化数据库
		2. 注册服务 etcd
	 */
	/* reg :=etcdv3.NewRegistry()  go run main.go –registrer=etcdv3 –registrer-address=http://192.168.1.110:4567 */

	//注册etcd
	//reg := etcdv3.NewRegistry(func(op *registry.Options){
	//	op.Addrs = []string{
	//		"http://127.0.0.1:32779",
	//	}
	//})
	reg :=etcdv3.NewRegistry()

	//创建服务
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.book"),
		micro.Version("latest"),
		micro.Address(":9984"),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*15),
	)

	//初始化服务
	service.Init()

	//将实现服务端的 API 注册到服务端
	if err:=book.RegisterGreeterHandler(service.Server(), new(handler.Book));err!=nil{
		log.Fatal(err)
	}

	//运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
