package handler

import (
	"context"
	"devices/category/domain/model"
	"devices/category/domain/service"
	categories "devices/category/proto"
)

type Category struct{
	CategoryService service.ICategoryService
}
func(c *Category)CreateCategory(ctx context.Context,request *categories.Create_Category_Request,response *categories.Create_Category_Response) error{
	Category := &model.Category{
		CategoryName: request.CategoryName,
		CategoryImage: request.CategoryImage,
		CategoryDescription: request.CategoryDescription,
		CategoryLevel: request.CategoryLevel,
		CategoryParent: request.CategoryParent,
	}
	_,err := c.CategoryService.AddCategory(Category)
	if err!=nil{
		return err
	}
	response.Message = "添加成功"
	return nil
}
func(c *Category)DeleteCategory(ctx context.Context,request *categories.Delete_Category_Request,response *categories.Delete_Category_Response) error{
	err := c.CategoryService.DeleteCategoryByID(request.CategoryId)
	if err!=nil{
		return err
	}
	response.Message = "删除成功"
	return nil
}
func(c *Category)FindCategoryById(ctx context.Context,request *categories.FindCateGoryById_Request,response *categories.Category_Response) error{
	category,err := c.CategoryService.FindCategoryByID(request.Id)
	if err != nil{
		return err
	}
	response.CategoryId = category.ID
	response.CategoryParent = category.CategoryParent
	response.CategoryLevel = category.CategoryLevel
	response.CategoryName = category.CategoryName
	response.CategoryDescription = category.CategoryDescription
	response.CategoryImage = category.CategoryImage
	return nil
}
func(c *Category)FindCategoryByName(ctx context.Context,request *categories.Find_CategoryByName_Request,response *categories.Category_Response) error {
	return nil
}
func(c *Category)FindCategoryByLevel(ctx context.Context,request *categories.Find_CategoryByLevel_Request,response *categories.Category_Response) error{
	return nil
}
func(c *Category)FindAllCategory(ctx context.Context,request *categories.Find_All_Request,response *categories.Find_All_Response) error{
	return nil
}
func(c *Category)UpdateCategory(ctx context.Context,request *categories.Create_Category_Request,response *categories.Update_Category_Response)error{
	Category := &model.Category{
		CategoryName: request.CategoryName,
		CategoryImage: request.CategoryImage,
		CategoryDescription: request.CategoryDescription,
		CategoryLevel: request.CategoryLevel,
		CategoryParent: request.CategoryParent,
	}
	err := c.CategoryService.UpdateCategory(Category)
	if err != nil{
		return err
	}
	response.Message = "更新成功"
	return nil
}