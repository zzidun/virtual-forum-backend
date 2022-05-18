package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
)

func CategoryList() (categoryList []*gin.H, err error) {
	categorys, err := dao.CategoryQuery()
	if err != nil {
		return
	}

	for _, category := range categorys {

		var userName string
		userId, err := dao.CategoryerQueryByCategoryId(category.ID)
		if err != nil {
			break
		}

		if userId != 0 {
			user, err := dao.UserQueryById(userId)
			if err != nil {
				break
			}
			userName = user.Name
		}

		categoryList = append(categoryList, &gin.H{
			"id":         fmt.Sprintf("%d", category.ID),
			"name":       category.Name,
			"speak":      fmt.Sprintf("%d", category.Speak),
			"follow":     fmt.Sprintf("%d", category.Follow),
			"categoryer": userName,
		})
	}

	return
}
