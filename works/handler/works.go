package handler

import (
	"context"
	"devices/common"
	"devices/works/domain/model"
	work "devices/works/domain/server"
	works "devices/works/proto"
)

type Works struct{
	WorkService work.IWorkerServices
}
// Call is a single request handler called via client.Call or the generated client code
func(w *Works)CreateWorker(ctx context.Context,req *works.Request_Workers,res *works.Response_CreateWorker)error{
	workers := &model.Workers{}
	err := common.SwapTo(req,workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	id,err := w.WorkService.CreateWorker(workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	res.Id = id
	return nil
}
func(w *Works)UpdateWorker(ctx context.Context,req *works.Request_Workers,res *works.Response_CreateWorker)error{
	workers := &model.Workers{}
	err := common.SwapTo(req,workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	id,err := w.WorkService.CreateWorker(workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	res.Id = id
	return nil
}
func(w *Works)DeleteWorkerByID(ctx context.Context,req *works.Request_Workers_ID,res *works.Response_Workers)error{
	err := w.WorkService.DeleteWorkerByID(req.Id)
	if err!=nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	return nil
}
func(w *Works)FindWorkerByID(ctx context.Context,req *works.Request_Workers_ID,res *works.Response_Worker_Show)error{
	worker,err := w.WorkService.FindWorkerByID(req.Id)
	if err!=nil{
		return err
	}
	err = common.SwapTo(res.Worker, worker)
	if err != nil{
		return err
	}
	return nil
}
func(w *Works)FindWorkerByName(ctx context.Context,req *works.Request_Workers_Name,res *works.Response_Workers_Show)error{
	worker,err := w.WorkService.FindWorkersByName(req.Name)
	if err != nil{
		return err
	}
	err = common.SwapTo(res.Workers, worker)
	if err != nil{
		return err
	}
	return nil
}
func(w *Works)FindAll(ctx context.Context,req *works.Request_Null,res *works.Response_Workers_Show)error{
	worker,err := w.WorkService.FindAll()
	if err != nil{
		return err
	}
	err = common.SwapTo(res.Workers, worker)
	if err != nil{
		return err
	}
	return nil
}