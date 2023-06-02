package model

import "gorm.io/gorm"

type RecordBasic struct {
	gorm.Model
	ProblemID int    `gorm:"column:problem_id; type: int;" json:"problem_id"`
	UserID    int    `gorm:"column:user_id; type: int;" json:"user_id"`
	Code      string `gorm:"column:code; type: varchar(255);" json:"code"`
	CodePath  string `gorm:"column:code_path; type: varchar(255);" json:"code_path"`
	language  string `gorm:"column:language; type: varchar(255);" json:"language"`
	status    int    `gorm:"column:status; type: int(4);" json:"status"`
}

func (table *RecordBasic) TableName() string {
	return "record_basic"
}

func GetRecordList(userId, problemId int, language string) *gorm.DB {
	database := DB
	if userId != -1 {
		database = database.Model(new(RecordBasic)).Where("user_id = ?", userId)
	}
	if problemId != -1 {
		database = database.Model(new(RecordBasic)).Where("problem_id = ?", problemId)
	}
	database = database.Model(new(RecordBasic)).Where("language = ?", language)

	return database
}
