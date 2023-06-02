package model

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name      string `gorm:"column:name; type:varchar(100);" json:"name"`
	Email     string `gorm:"column:email; type:varchar(100);" json:"email"`
	Password  string `gorm:"column:password; type:varchar(100);" json:"password"`
	Phone     string `gorm:"column:phone; type:varchar(100);" json:"phone"`
	ACNum     int    `gorm:"column:ac_num; type:int;" json:"ac_num"`
	SubmitNum int    `gorm:"column:submit_num; type:int;" json:"submit_num"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
