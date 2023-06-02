package model

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemID     int            `gorm:"column:problem_id; type:int;" json:"problem_id"`
	CategoryID    int            `gorm:"column:category_id; type:int;" json:"category_id"`
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id; reference:category_id"`
}

func (Table ProblemCategory) TableName() string {
	return "problem_category"
}
