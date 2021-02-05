package main

import (
	"device/handler"
	pb "device/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("device"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterDeviceHandler(srv.Server(), new(handler.Device))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
