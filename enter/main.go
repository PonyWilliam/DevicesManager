package main

import (
	"enter/handler"
	pb "enter/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("enter"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterEnterHandler(srv.Server(), new(handler.Enter))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
