package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
)

func CategoryList(left int, right int) (categoryList *gin.H, err error) {
	categorys, totNum, curNum, err := dao.CategoryQuery(left, right)
	if err != nil {
		return
	}

	var categoryListData []*gin.H

	for _, category := range categorys {

		var userName string
		categoryer, err := dao.CategoryerQueryByCategoryId(category.ID)
		if err != nil {
			continue
		}

		if categoryer.ID != 0 {
			user, err := dao.UserQueryById(categoryer.ID)
			if err != nil {
				continue
			}
			userName = user.Name
		}

		categoryListData = append(categoryListData, &gin.H{
			"id":         fmt.Sprintf("%d", category.ID),
			"name":       category.Name,
			"speak":      fmt.Sprintf("%d", category.Speak),
			"follow":     fmt.Sprintf("%d", category.Follow),
			"categoryer": userName,
		})
	}

	categoryList = &gin.H{
		"tot":  fmt.Sprintf("%d", totNum),
		"cur":  fmt.Sprintf("%d", curNum),
		"list": categoryListData,
	}

	return
}

func CategoryerCheck(userId uint, categoryId uint) (ret bool, err error) {
	categoryerId, err := dao.CategoryerQueryByCategoryId(categoryId)

	ret = (categoryerId.ID == userId)

	return
}
