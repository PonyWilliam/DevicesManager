package model
type WorkLog struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	WorkID int64 `json:"work_id"`//员工id
	ProductID int64 `json:"product_id"`
	Description string `json:"description"`//用途描述(一键借走则可空)
}