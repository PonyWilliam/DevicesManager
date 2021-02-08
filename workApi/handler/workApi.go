package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	works "github.com/PonyWilliam/go-works/proto"
	log "github.com/micro/go-micro/v2/logger"
	"strconv"
	workApi "workApi/proto/workApi"
)

type WorkApi struct {
	WorksServices works.WorksService
}
func(e *WorkApi) FindAll(ctx context.Context,req *workApi.Request,res *workApi.Response) error {
	log.Info("接收到返回请求")
	if _,ok := req.Get["user_id"];!ok{
		res.StatusCode = 500
		return errors.New("参数异常")
	}
	userIdString := req.Get["user_id"].Value[0]
	fmt.Println(userIdString)
	userId,err := strconv.ParseInt(userIdString,10,64)
	if err != nil{
		return err
	}
	//获取用户信息表
	WorkAll,err := e.WorksServices.FindWorkerByID(context.TODO(),&works.Request_Workers_ID{Id: userId})//不需要传入
	b,err := json.Marshal(WorkAll)
	if err!= nil{
		return err
	}
	res.StatusCode = 200
	res.Body = string(b)
	return nil
}