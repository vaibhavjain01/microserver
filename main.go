package main

import (
	"hello/handler"
	pb "hello/proto"
	"time"
	"context"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	log "github.com/micro/micro/v3/service/logger"
)

func client() {
	// create and initialise a new service
	srv := service.New()

	// create the proto client for helloworld
	client := pb.NewHelloService("hello", srv.Client())

	// call an endpoint on the service
	rsp, err := client.Call(context.Background(), &pb.Request{
		Name: "John",
	})
	if err != nil {
		log.Info(err)
		return
	}

	// print the response
	log.Info("Response: " + rsp.Msg)
	
	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
}

func main() {
	
	// Create service
	srv := service.New(
		service.Name("hello"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterHelloHandler(srv.Server(), new(handler.Hello))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
