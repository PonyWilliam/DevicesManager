package handler

import (
	"context"
	"devices/common"
	"devices/works/domain/model"
	work "devices/works/domain/server"
	works "devices/works/proto"
)

type Works struct{
	workService work.WorkServices
}

// Call is a single request handler called via client.Call or the generated client code
func(w *Works)CreateWorker(ctx context.Context,req *works.Request_Workers,res *works.Response_CreateWorker)error{
	workers := &model.Workers{}
	err := common.SwapTo(req,workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	id,err := w.workService.CreateWorker(workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	res.Id = id
	return nil
}