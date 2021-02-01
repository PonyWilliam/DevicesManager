package main

import (
	"devices/common"
	"devices/product/domain/repository"
	services2 "devices/product/domain/service"
	"devices/product/handler"
	product2 "devices/product/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"strconv"
	"time"
)

func main() {
	consulConfig,err := common.GetConsualConfig("127.0.0.1",8500,"/micro/config")
	//配置中心
	if err != nil{
		log.Fatal(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(
		func(options *registry.Options){
			options.Addrs = []string{"127.0.0.1"}
			options.Timeout = time.Second * 10
		},
		)

	srv := micro.NewService(
		micro.Name("services.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(consulRegistry),
		)
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	db,err := gorm.Open("mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host + ":"+ strconv.FormatInt(mysqlInfo.Port,10) +")/"+mysqlInfo.DataBase+"?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil{
		log.Error(err)

	}
	defer db.Close()
	db.SingularTable(true)
	srv.Init()
	rp := repository.NewProductRepository(db)
	rp.InitTable()

	productServices := services2.NewProductServices(repository.NewProductRepository(db))
	err = product2.RegisterProductHandler(srv.Server(),&handler.Product{ProductServices: productServices})

	if err:=srv.Run();err!=nil{
		log.Fatal(err)
	}
}