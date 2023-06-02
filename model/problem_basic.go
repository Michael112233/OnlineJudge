package model

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Title           string           `gorm:"column:title; type:varchar(255);" json:"title"`  //标题
	Content         string           `gorm:"column:content; type:text;" json:"content"`      //内容
	PassNum         int              `gorm:"column:pass_num; type:int;" json:"pass_num"`     //通过人数
	SubmitNum       int              `gorm:"column:submit_num; type:int;" json:"submit_num"` //提交人数
	MaxTime         int              `gorm:"column:maxTime; type:int;" json:"max_time"`
	MaxMem          int              `gorm:"column:maxMem; type:int;" json:"max_mem"`
	ProblemCategory *ProblemCategory `gorm:"foreignKey:problem_id; reference:id"`
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, category string) *gorm.DB {
	selectDB := DB.Model(new(ProblemBasic)).Where("title like ? or content like ?", "%"+keyword+"%", "%"+keyword+"%")

	var cid int
	if category != "" {
		selectDB.Raw("select id from category_basic where name = ?", category).Scan(&cid)
		//fmt.Println(cid)
		selectDB.Joins("right join problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = ?", cid)
	}

	return selectDB
}
