package main

import (
	"context"
	"fmt"
	"github.com/phiskills/grpc-api.go"
	"github.com/phisuite/data.go"
	"log"
)

type entityServer struct {
	data.UnimplementedEntityAPIServer
}

func (e entityServer) List(_ *data.Options, stream data.EntityAPI_ListServer) error {
	for i := 1; i < 5; i++ {
		version := fmt.Sprintf("0.0.%d", i)
		entity := &data.Entity{Name:"dummy", Version:version}
		log.Printf("Stream: %v", entity)
		if err := stream.Send(entity); err != nil {
			return err
		}
	}
	return nil
}

func (e entityServer) Get(context.Context, *data.Options) (*data.Entity, error) {
	entity := &data.Entity{Name:"dummy", Version:"0.0.1"}
	log.Printf("Send: %v", entity)
	return entity, nil
}

func main() {
	api := grpc.New()
	data.RegisterEntityAPIServer(api.Server, &entityServer{})
	api.Start()
}
