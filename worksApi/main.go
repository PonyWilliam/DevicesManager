package main

import (
	"context"
	"fmt"
	go_worker_services "github.com/PonyWilliam/go-works/proto"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	"net"
	"net/http"
	"time"
	"worksApi/handler"
	worksApi "worksApi/proto"
)

func main() {
	consul := consul2.NewRegistry(
		func(options *registry.Options) {
			options.Timeout = time.Second * 10
			options.Addrs = []string{"127.0.0.1:8500"}
		},
	)
	hystrixStreamHandler := hystrix.NewStreamHandler()//创建熔断器
	hystrixStreamHandler.Start()
	 go func() {
	 	err := http.ListenAndServe(net.JoinHostPort("0.0.0.0","9096"),hystrixStreamHandler)
	 	if err!=nil{
	 		log.Error(err)
		}
	 }()
	services := micro.NewService(
		micro.Name("go.micro.api.workApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8085"),
		micro.Registry(consul),
		//添加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		//添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),

	)
	services.Init()
	workerServices := go_worker_services.NewWorksService("go.micro.service.works",services.Client())
	if err:=worksApi.RegisterWorkApiHandler(services.Server(),&handler.WorkApi{WorksServices: workerServices});err!=nil{
		log.Error(err)
	}


	if err:= services.Run();err!=nil{
		log.Error(err)
	}
}
type clientWrapper struct{
	client.Client
}
func (c *clientWrapper) Call(ctx context.Context,req client.Request,res interface{},opts ...client.CallOption) error {
	return hystrix.Do(req.Service() + "." + req.Endpoint(),
		//正常执行逻辑
		func() error {
			return c.Client.Call(ctx,req,res,opts...)
		},//错误逻辑
		func(err error) error {
			fmt.Println(err)
			return err
		},
	)
}

func NewClientHystrixWrapper() client.Wrapper{
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}