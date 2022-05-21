package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
)

func CommentList(postId uint, left int, right int) (commentList *gin.H, err error) {
	comments, totNum, curNum, err := dao.CommentQueryByPostId(postId, left, right)
	if err != nil {
		return
	}

	post, err := dao.PostQueryById(postId)
	if err != nil {
		return
	}

	var commentListData []*gin.H
	for _, comment := range comments {

		user, err := dao.UserQueryById(comment.UserId)
		if err != nil {
			continue
		}

		userFollow, err := dao.UserFollowQuery(user.ID, post.CategoryId)
		if err != nil {
			continue
		}

		commentListData = append(commentListData, &gin.H{
			"id":      fmt.Sprintf("%d", comment.ID),
			"replyid": fmt.Sprintf("%d", comment.ReplyId),
			"content": comment.Content,
			"ctime":   comment.CreatedAt,
			"user": &gin.H{
				"id":    fmt.Sprintf("%d", user.ID),
				"name":  user.Name,
				"speak": fmt.Sprintf("%d", user.Speak),
				"count": fmt.Sprintf("%d", userFollow.Count),
			},
		})
	}

	commentList = &gin.H{
		"tot":  fmt.Sprintf("%d", totNum),
		"cur":  fmt.Sprintf("%d", curNum),
		"list": commentListData,
	}

	return
}
