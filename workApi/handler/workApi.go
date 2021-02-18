package handler
/*
	rpc CreateWorker(Request)	returns(Response);
	rpc UpdateWorker(Request)	returns(Response);
	rpc DeleteWorkerByID(Request)	returns(Response);
	rpc FindWorkerByID(Request)	returns(Response);
	rpc FindWorkerByName(Request)	returns(Response);

	rpc CreateToken(Request)	returns(Response);
	rpc FindAll(Request)	returns(Response);
*/
import (
	"context"
	"encoding/json"
	"fmt"
	works "github.com/PonyWilliam/go-works/proto"
	"github.com/micro/go-micro/v2/util/log"
	"strconv"
	workApi "workApi/proto/workApi"
)

type Workers struct{
	Client works.WorksService
}
type Response struct{
	id int64
	msg string
}
func(w *Workers)CreateWorker(ctx context.Context,req *workApi.Request,rsp *workApi.Response)error{
	Name := req.Get["name"].Value[0]
	Nums := req.Get["nums"].Value[0]
	Sex := req.Get["sex"].Value[0]
	Level := req.Get["level"].Value[0]
	Score := req.Get["score"].Value[0]
	Place := req.Get["place"].Value[0]
	Telephone := req.Get["telephone"].Value[0]
	Description := req.Get["description"].Value[0]
	ISWork := req.Get["iswork"].Value[0]
	Username := req.Get["username"].Value[0]
	Password := req.Get["password"].Value[0]
	fmt.Println(Username)
	Level_int, err := strconv.ParseInt(Level, 10, 64)
	if err!=nil{
		rsp.Body = "Level转换错误"
		rsp.StatusCode = 500
		return err
	}
	Score_int, err := strconv.ParseInt(Score, 10, 64)
	if err!=nil{
		rsp.Body = "Level转换错误"
		rsp.StatusCode = 500
		return err
	}
	Is_Work, err := strconv.ParseBool(ISWork)
	if err!=nil{
		rsp.Body = "Level转换错误"
		rsp.StatusCode = 500
		return err
	}
	worker,err := w.Client.CreateWorker(context.TODO(),&works.Request_Workers{Name: Name,
		Nums: Nums,
		Sex: Sex,
		Level:Level_int,
		Score: Score_int,
		Place: Place,
		Telephone: Telephone,
		Description: Description,
		Username: Username,
		Password: Password,
		ISWork: Is_Work,
	})
	if err!=nil{
		rsp.StatusCode = 500
		rsp.Body = "创建失败"
	}
	fmt.Println(worker)
	b,err := json.Marshal(&Response{1,"创建成功"})
	if err!=nil{
		rsp.Body = "转换失败"
		rsp.StatusCode = 500
		return err
	}

	rsp.Body = string(b)
	rsp.StatusCode = 200
	return nil
}
func(w *Workers)UpdateWorker(ctx context.Context,req *workApi.Request,rsp *workApi.Response)error{
	Name := req.Get["name"].Value[0]
	Nums := req.Get["nums"].Value[0]
	Sex := req.Get["sex"].Value[0]
	Level := req.Get["level"].Value[0]
	Score := req.Get["score"].Value[0]
	Place := req.Get["place"].Value[0]
	Telephone := req.Get["telephone"].Value[0]
	Description := req.Get["description"].Value[0]
	ISWork := req.Get["iswork"].Value[0]
	Username := req.Get["username"].Value[0]
	Password := req.Get["password"].Value[0]
	Level_int, err := strconv.ParseInt(Level, 10, 64)
	if err!=nil{
		rsp.Body = "Level转换错误"
		rsp.StatusCode = 500
		return err
	}
	Score_int, err := strconv.ParseInt(Score, 10, 64)
	if err!=nil{
		rsp.Body = "Level转换错误"
		rsp.StatusCode = 500
		return err
	}
	Is_Work, err := strconv.ParseBool(ISWork)
	if err!=nil{
		rsp.Body = "Level转换错误"
		rsp.StatusCode = 500
		return err
	}
	worker,err := w.Client.CreateWorker(context.TODO(),&works.Request_Workers{Name: Name,
		Nums: Nums,
		Sex: Sex,
		Level:Level_int,
		Score: Score_int,
		Place: Place,
		Telephone: Telephone,
		Description: Description,
		Username: Username,
		Password: Password,
		ISWork: Is_Work,
	})
	if err!=nil{
		rsp.StatusCode = 500
		rsp.Body = "创建失败"
	}
	b,err := json.Marshal(&Response{worker.Id,"创建成功"})
	if err!=nil{
		rsp.Body = "转换失败"
		rsp.StatusCode = 500
		return err
	}

	rsp.Body = string(b)
	rsp.StatusCode = 200
	return nil
}
func(w *Workers)FindWorkerByID(ctx context.Context,req *workApi.Request,rsp *workApi.Response)error{

	id := req.Get["id"].Value[0]
	user_id,err := strconv.ParseInt(id,10,64)
	if err!= nil{
		rsp.StatusCode = 500
		rsp.Body = err.Error()
		return err
	}
	worker,err := w.Client.FindWorkerByID(context.TODO(),&works.Request_Workers_ID{Id: user_id})
	if err!=nil{
		rsp.StatusCode = 500
		rsp.Body = "查找错误"
		return err
	}
	b,err := json.Marshal(worker)
	rsp.Body = string(b)
	rsp.StatusCode = 200
	return nil
}
func(w *Workers)FindWorkerByNums(ctx context.Context,req *workApi.Request,rsp *workApi.Response) error{
	nums := req.Get["id"].Value[0]
	nums_int,err := strconv.ParseInt(nums,10,64)
	if err!= nil{
		rsp.StatusCode = 500
		rsp.Body = err.Error()
		return err
	}
	worker,err := w.Client.FindWorkerByNums(context.TODO(),&works.Request_Workers_Nums{
		Nums: nums_int,
	})
	if err!=nil{
		rsp.StatusCode = 500
		rsp.Body = "查找错误"
		return err
	}
	b,err := json.Marshal(worker)
	rsp.Body = string(b)
	rsp.StatusCode = 200
	return nil
}
func(w *Workers)DeleteWorkerByID(ctx context.Context,req *workApi.Request,rsp *workApi.Response) error {
	id := req.Get["id"].Value[0]
	user_id,err := strconv.ParseInt(id,10,64)
	if err != nil{
		rsp.StatusCode = 500
		rsp.Body = err.Error()
		return err
	}
	res,err := w.Client.DeleteWorkerByID(context.TODO(),&works.Request_Workers_ID{
		Id: user_id,
	})
	if err != nil{
		rsp.StatusCode = 500
		rsp.Body = err.Error()
	}
	rsp.StatusCode = 200
	rsp.Body = res.Message
	return nil
}
func(w *Workers)FindWorkerByName(ctx context.Context,req *workApi.Request,rsp *workApi.Response)error{
	name := req.Get["name"].Value[0]
	workers,err := w.Client.FindWorkerByName(context.TODO(),&works.Request_Workers_Name{Name: name})
	if err!=nil{
		rsp.StatusCode = 500
		rsp.Body = "查找错误"
		log.Error(err)
		return err
	}
	b,err := json.Marshal(workers)
	if err!=nil{
		rsp.Body = "转换json错误"
		rsp.StatusCode = 500
	}
	rsp.Body = string(b)
	rsp.StatusCode = 200

	return nil
}
func(w *Workers)FindAll(ctx context.Context,req *workApi.Request,rsp *workApi.Response)error{
	workers,err := w.Client.FindAll(context.TODO(),&works.Request_Null{})
	if err!=nil{
		rsp.StatusCode = 500
		rsp.Body = "查找错误"
		log.Error(err)
		return err
	}
	b,err := json.Marshal(workers)
	if err!=nil{
		rsp.Body = "转换json错误"
		rsp.StatusCode = 500
	}
	rsp.Body = string(b)
	rsp.StatusCode = 200

	return nil
}