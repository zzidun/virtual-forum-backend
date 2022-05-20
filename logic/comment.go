package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
)

func CommentList(postId uint, left int, right int) (commentList []*gin.H, err error) {
	comments, err := dao.CommentQueryByPostId(postId, left, right)
	if err != nil {
		return
	}

	for _, comment := range comments {

		user, err := dao.UserQueryById(comment.UserId)
		if err != nil {
			break
		}

		commentList = append(commentList, &gin.H{
			"id":       fmt.Sprintf("%d", comment.ID),
			"username": user.Name,
			"replyid":  fmt.Sprintf("%d", comment.ReplyId),
			"content":  comment.Content,
		})
	}

	return
}
