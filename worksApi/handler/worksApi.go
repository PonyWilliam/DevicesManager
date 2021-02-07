package handler

import (
	"context"
	works "devices/works/proto"
	worksApi "devices/worksApi/proto"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/micro/micro/v3/service/logger"
	"strconv"
)

type WorksApi struct {
	WorksServices works.WorksService
}
func(e *WorksApi) FindAll(ctx context.Context,req *worksApi.Request,res *worksApi.Response) error {
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