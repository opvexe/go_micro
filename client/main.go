package main

import(
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pbw "ibook/server/book/proto/book"
)

func main() {
	//注册etcd
	/* reg :=etcdv3.NewRegistry() */
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:32779",
		}
	})

	sevice := micro.NewService(
		micro.Registry(reg),
	)
	sevice.Init()
	c := pbw.NewGreeterService("go.micro.srv.book",sevice.Client())
	resp,err:=c.SayHello(context.Background(),&pbw.HelloRequest{
		Name:                 "SHUMIN",
	})
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Message)
}
