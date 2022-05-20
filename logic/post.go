package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
)

func PostList(categoryId uint, left int, right int) (PostList []*gin.H, err error) {
	posts, err := dao.PostQueryByCategoryId(categoryId, left, right)
	if err != nil {
		return
	}

	for _, post := range posts {

		user, err := dao.UserQueryById(post.UserId)
		if err != nil {
			break
		}

		PostList = append(PostList, &gin.H{
			"id":    fmt.Sprintf("%d", post.ID),
			"title": post.Title,
			"speak": fmt.Sprintf("%d", post.Speak),
			"user":  user.Name,
		})
	}

	return
}
