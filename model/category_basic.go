package model

import "gorm.io/gorm"

type CategoryBasic struct {
	gorm.Model
	Name            string           `gorm:"column:name; type:varchar(255);" json:"name"`
	ParentID        int              `gorm:"column:parent_id; type:int;" json:"parent_id"`
	ProblemCategory *ProblemCategory `gorm:"foreignKey:category_id; reference:id"`
}

func (table *CategoryBasic) TableName() string {
	return "category_basic"
}
