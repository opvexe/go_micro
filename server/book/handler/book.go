package handler

import (
	"context"
	book "ibook/server/book/proto/book"
)

type Book struct{}


func (e *Book) SayHello(ctx context.Context, req *book.HelloRequest, rsp *book.HelloReply) error {
	//返回测试数据
	rsp.Message = "我是测试数据"
	return nil
}