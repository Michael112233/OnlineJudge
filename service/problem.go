package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"onlineJudge/define"
	"onlineJudge/model"
	"strconv"
)

// GetProblemList
// @Summary Problem List
// @Description get Problem List
// @Tags Public implement
// @Param size query int false "size"
// @Param page query int false "page"
// @Param keyword query string false "keyword"
// @Param category query string false "category"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem [get]
func GetProblemList(txt *gin.Context) {
	// get page and size number
	page := txt.DefaultQuery("page", define.DefaultPage)
	size := txt.DefaultQuery("size", define.DefaultSize)

	pageNum, errPage := strconv.Atoi(page)
	sizeNum, errSize := strconv.Atoi(size)
	if errPage != nil {
		log.Println("Page Input Error!")
	}
	if errSize != nil {
		log.Println("Size Input Error!")
	}
	keyword := txt.Query("keyword")
	category := txt.Query("category")

	var count int64 = 0
	db := model.GetProblemList(keyword, category)
	data := make([]*model.ProblemBasic, 0)
	err := db.Count(&count).Omit("content").Offset((pageNum - 1) * sizeNum).Limit(sizeNum).Find(&data).Error
	if err != nil {
		log.Println("Problem Not Found!")
	}
	//txt.String(http.StatusOK, "Get Problem List")

	txt.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "find OK!",
		//"data":  data,
		//"count": count,
		"data": map[string]interface{}{
			"data":  data,
			"count": count,
		},
	})
}

// GetProblemDetail
// @Summary Problem Detail
// @Description get Problem Detail
// @Tags Public implement
// @Param problemID query int false "problemID"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem_detail [get]
func GetProblemDetail(c *gin.Context) {
	problemIdStr := c.Query("problemID")
	if problemIdStr == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Problem ID can not be null",
		})
		return
	}
	problemId, err := strconv.Atoi(problemIdStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Problem ID not be an int",
		})
		return
	}
	fmt.Println("---------")
	fmt.Println(problemId)

	data := new(model.ProblemBasic)
	err = model.DB.Where("id = ?", problemId).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "Problem not found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "Other errors!" + err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Successfully found!",
		"data": data,
	})
}
