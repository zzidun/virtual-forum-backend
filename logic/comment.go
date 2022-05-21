package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/dao"
)

func CommentList(postId uint, left int, right int, userId uint) (commentList *gin.H, err error) {

	// 查询一页16个评论
	comments, totNum, curNum, err := dao.CommentQueryByPostId(postId, left, right)
	if err != nil {
		return
	}

	post, err := dao.PostQueryById(postId)
	if err != nil {
		return
	}

	userShields, err := dao.UserShieldQueryByUser1(userId)
	if err != nil {
		return
	}

	var commentListData []*gin.H
	for _, comment := range comments {

		flag := false
		for _, userShield := range userShields {
			if userShield.ShieldUserId == comment.UserId {
				flag = true
				break
			}
		}

		if flag {
			continue
		}

		// 添加评论到返回结果

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
