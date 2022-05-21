package logic

import (
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

func AdminList(left int, right int) (adminList []*gin.H, err error) {
	return
}
