package main

import(
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb "ibook/server/book/proto/book"
)

func main() {

	//注册etcd
	/* reg :=etcdv3.NewRegistry() */
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://172.19.0.4:2379",
		}
	})

	sevice := micro.NewService(
		micro.Registry(reg),
	)
	client := pb.NewGreeterService("go.micro.srv.book",sevice.Client())
	resp,err:=client.SayHello(context.Background(),&pb.HelloRequest{
		Name:                 "SHUMIN",
	})
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Message)
}
