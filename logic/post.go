package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
)

func PostList(categoryId uint, left int, right int) (postList *gin.H, err error) {
	posts, totNum, curNum, err := dao.PostQueryByCategoryId(categoryId, left, right)
	if err != nil {
		return
	}

	var postListData []*gin.H
	for _, post := range posts {

		user, err := dao.UserQueryById(post.UserId)
		if err != nil {
			continue
		}

		userFollow, err := dao.UserFollowQuery(user.ID, post.CategoryId)
		if err != nil {
			continue
		}

		postListData = append(postListData, &gin.H{
			"id":    fmt.Sprintf("%d", post.ID),
			"title": post.Title,
			"speak": fmt.Sprintf("%d", post.Speak),
			"user": &gin.H{
				"id":    fmt.Sprintf("%d", user.ID),
				"name":  user.Name,
				"speak": fmt.Sprintf("%d", user.Speak),
				"count": fmt.Sprintf("%d", userFollow.Count),
			},
		})
	}

	postList = &gin.H{
		"tot":  fmt.Sprintf("%d", totNum),
		"cur":  fmt.Sprintf("%d", curNum),
		"list": postListData,
	}

	return
}
