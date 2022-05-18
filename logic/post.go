package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
)

func PostList(categoryId uint) (PostList []*gin.H, err error) {
	posts, err := dao.PostQueryByCategoryId(categoryId)
	if err != nil {
		return
	}

	for _, post := range posts {

		user, err := dao.UserQueryById(post.UserId)
		if err != nil {
			break
		}

		PostList = append(PostList, &gin.H{
			"id":         fmt.Sprintf("%d", post.ID),
			"title":      post.Title,
			"speak":      fmt.Sprintf("%d", post.Speak),
			"categoryer": user.Name,
		})
	}

	return
}
