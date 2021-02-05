package model
type WorkLog struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	WorkID int64 `json:"work_id"`//员工id

}