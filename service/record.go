package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineJudge/define"
	"onlineJudge/model"
	"strconv"
)

// GetRecordDetail
// @Summary Record List
// @Description get Record List
// @Tags Public implement
// @Param problemID query int false "problem_id"
// @Param userID query int false "user_id"
// @Param language query string false "language"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /record_detail [get]
func GetRecordDetail(c *gin.Context) {
	problemIdStr := c.Query("problemID")
	fmt.Println("1" + problemIdStr)
	if problemIdStr == "" {
		problemIdStr = "-1"
	}
	problemId, err := strconv.Atoi(problemIdStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "problem_id error!",
		})
		return
	}

	userIdStr := c.Query("userID")
	fmt.Println("2" + userIdStr)
	if userIdStr == "" {
		userIdStr = "-1"
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "user_id error!",
		})
		return
	}

	language := c.DefaultQuery("language", define.DefaultSubmitLanguage)

	selectDB := model.GetRecordList(userId, problemId, language)
	data := make([]*model.RecordBasic, 0)
	err = selectDB.Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "successfully found",
		"data": data,
	})
}
