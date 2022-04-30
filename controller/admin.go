package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/model"
)

func AdminLogin(c *gin.Context) {
	var admin_form *model.AdminLoginForm
	if err := c.ShouldBindJSON(&admin_form); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}

	admin := &model.Admin{
		Name: admin_form.Name,
		Password: admin_form.Password,
	}
	
	if err := dao.AdminLogin(admin); err != nil {
		if errors.Is(err,mysql.ErrorUserNotExit) {
			ResponseError(c,CodeUserNotExist)
			return
		}
		ResponseError(c,CodeInvalidParams)
		return
	}

	ResponseSuccess(c, gin.H{
		"admin_id" : fmt.Sprintf("%d", admin.ID),
		"admin_name": admin.Name,
		"access_token":,
		"refresh_token":
	})
}

func RefreshTokenHandler(c *gin.Context) {

}

