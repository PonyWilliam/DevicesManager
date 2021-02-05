package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	works "works/proto"
)

type Works struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Works) Call(ctx context.Context, req *works.Request, rsp *works.Response) error {
	log.Info("Received Works.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Works) Stream(ctx context.Context, req *works.StreamingRequest, stream works.Works_StreamStream) error {
	log.Infof("Received Works.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&works.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Works) PingPong(ctx context.Context, stream works.Works_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&works.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
