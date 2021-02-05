package main

import (
	"manager/handler"
	pb "manager/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("manager"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterManagerHandler(srv.Server(), new(handler.Manager))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
