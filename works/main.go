package main

import (
	"works/handler"
	pb "works/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("works"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterWorksHandler(srv.Server(), new(handler.Works))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
