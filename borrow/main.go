package main

import (
	"borrow/handler"
	pb "borrow/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("borrow"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterBorrowHandler(srv.Server(), new(handler.Borrow))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
