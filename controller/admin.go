package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
)

func AdminLogin(c *gin.Context) {
	var admin_form *model.AdminLoginForm
	if err := c.ShouldBindJSON(&admin_form); err != nil {
		response.ResponseError(c, response.CodeInvalidParams)
		return
	}

	admin := &model.Admin{
		Name:     admin_form.Name,
		Password: admin_form.Password,
	}

	if err := dao.AdminLogin(admin); err != nil {
		// if errors.Is(err, mysql.ErrorUserNotExit) {
		// 	response.ResponseError(c, response.CodeUserNotExist)
		// 	return
		// }
		response.ResponseError(c, response.CodeInvalidParams)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"admin_id":      fmt.Sprintf("%d", admin.ID),
		"admin_name":    admin.Name,
		"access_token":  1,
		"refresh_token": 1,
	})
}

func RefreshTokenHandler(c *gin.Context) {

}
