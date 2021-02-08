package client

import (
	"context"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	workApi "path/to/service/proto/workApi"
)

type workApiKey struct {}

// FromContext retrieves the client from the Context
func WorkApiFromContext(ctx context.Context) (workApi.WorkApiService, bool) {
	c, ok := ctx.Value(workApiKey{}).(workApi.WorkApiService)
	return c, ok
}

// Client returns a wrapper for the WorkApiClient
func WorkApiWrapper(service micro.Service) server.HandlerWrapper {
	client := workApi.NewWorkApiService("go.micro.service.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, workApiKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
