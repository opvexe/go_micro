package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	book "ibook/server/book/proto/book"
)

type Book struct{}

func (e *Book) Handle(ctx context.Context, msg *book.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *book.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
