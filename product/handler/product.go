package handler

import (
	"context"
	"devices/product/domain/model"
	"devices/product/domain/service"
	products "devices/product/proto"
	"strconv"
)
/*
package model
type Product struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	ProductName string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	Level int64 `json:"level"`
	Category int64 `json:"category"`//指向categoryid
	Important bool `json:"important"`//说明是否重要
	Is bool `json:"is"`//是否在库
	BelongCustom int64 `json:"belong_custom"`//当前所属用户ID
	BelongArea int64 `json:"belong_area"`//所属库房
	Location string `json:"location"`//最新的定位信息
	Rfid string `json:"rfid"`//rfid标记
	ImageID int64 `json:"image_id"`//图片地址对应的id（可上传）
}

*/

type Product struct{
	ProductServices service.IProductServices
}
func(p *Product)AddProduct(ctx context.Context,info *products.Request_ProductInfo,product *products.Response_Product) error {
	IProduct := &model.Product{
		ID: info.Id,
		ProductName: info.ProductName,
		Level: info.ProductLevel,
		ProductDescription: info.ProductDescription,
		Category: info.ProductBelongCategory,
		Important: info.IsImportant,
		Is:info.ProductIs,
		BelongCustom: info.ProductBelongCustom,
		BelongArea: info.ProductBelongArea,
		Location: info.ProductLocation,
		Rfid: info.ProductRfid,
		ImageID:info.ImageId,
	}
	id,err := p.ProductServices.AddProduct(IProduct)
	if err!=nil{
		product.Message = err.Error()
		return err
	}
	product.Message = strconv.FormatInt(id, 10)
	return nil
}
func(p *Product)DelProduct(ctx context.Context,id *products.Request_ProductID,product *products.Response_DelProduct)error{
	err := p.ProductServices.DelProductByID(id.Id)
	if err!=nil{
		product.Message = err.Error()
		return err
	}
	return nil
}
func(p *Product)ChangeProduct(ctx context.Context,info *products.Request_ProductInfo,product *products.Response_Product)error{
	IProduct := &model.Product{
		ID: info.Id,
		ProductName: info.ProductName,
		Level: info.ProductLevel,
		ProductDescription: info.ProductDescription,
		Category: info.ProductBelongCategory,
		Important: info.IsImportant,
		Is:info.ProductIs,
		BelongCustom: info.ProductBelongCustom,
		BelongArea: info.ProductBelongArea,
		Location: info.ProductLocation,
		Rfid: info.ProductRfid,
		ImageID:info.ImageId,
	}
	id,err := p.ProductServices.AddProduct(IProduct)
	if err!=nil{
		product.Message = err.Error()
		return err
	}
	product.Message = strconv.FormatInt(id, 10)
	return nil
}
// Call is a single request handler called via client.Call or the generated client code
