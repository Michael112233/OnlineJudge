package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineJudge/model"
)

// GetUserDetail
// @Summary user Detail
// @Description get user detail
// @Tags Public implement
// @Param username query string false "username"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /user_detail [get]
func GetUserDetail(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "username cannot be null!",
		})
		return
	}

	data := make([]*model.UserBasic, 0)
	err := model.DB.Omit("password", "phone", "email").Where("name like ?", "%"+username+"%").Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "cannot find relative user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Successfully find",
		"data": data,
	})
}
