package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"onlineJudge/model"
	"testing"
)

func forTest(t *testing.T) {
	dsn := "root:123456@(localhost)/onlineJudge?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("Connection error!")
	}

	data := make([]*model.ProblemBasic, 0)

	err = db.Find(&data).Error
	if err != nil {
		t.Fatal("Find error!")
	}

	for _, v := range data {
		fmt.Printf("Problem -> %v\n", v)
	}
}
