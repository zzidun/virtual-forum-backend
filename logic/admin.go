package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
)

type AdminPermCode int64

const (
	CodeAdminPerm    AdminPermCode = 1000
	CodeBanPerm      AdminPermCode = 1001
	CodeCategoryPerm AdminPermCode = 1002
)

func UserCheckPerm(userId uint, code AdminPermCode) (valid uint, err error) {
	user, err := dao.UserQueryById(userId)
	if err != nil {
		return 0, err
	}

	admin, err := dao.AdminQueryById(user.AdminId)
	if err != nil {
		return 0, err
	}

	switch code {
	case CodeAdminPerm:
		valid = admin.AdminPerm
	case CodeBanPerm:
		valid = admin.BanPerm
	case CodeCategoryPerm:
		valid = admin.CategoryPerm
	default:
		return 0, nil
	}

	return
}

func AdminList(left int, right int) (adminList *gin.H, err error) {
	admins, totNum, curNum, err := dao.AdminQuery(left, right)
	if err != nil {
		return
	}

	var adminListData []*gin.H

	for _, admin := range admins {

		user, err := dao.UserQueryById(admin.UserId)
		if err != nil {
			continue
		}

		adminListData = append(adminListData, &gin.H{
			"userid":       fmt.Sprintf("%d", user.ID),
			"username":     user.Name,
			"adminid":      fmt.Sprintf("%d", admin.ID),
			"adminperm":    fmt.Sprintf("%d", admin.AdminPerm),
			"banperm":      fmt.Sprintf("%d", admin.BanPerm),
			"categoryperm": fmt.Sprintf("%d", admin.CategoryPerm),
		})
	}

	adminList = &gin.H{
		"tot":  fmt.Sprintf("%d", totNum),
		"cur":  fmt.Sprintf("%d", curNum),
		"list": adminListData,
	}

	return
}
