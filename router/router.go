package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "onlineJudge/docs"
	"onlineJudge/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	//make swaggers
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// for problem
	r.GET("/problem", service.GetProblemList)
	r.GET("/problem_detail", service.GetProblemDetail)
	// for user
	r.GET("/user_detail", service.GetUserDetail)
	// for record
	r.GET("/record_detail", service.GetRecordDetail)
	return r
}

// swag init 生成 docs 文件夹
