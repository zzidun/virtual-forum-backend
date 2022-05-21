package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

// 更新版块关注人数
func CategoryCountFollow(categoryId uint, dir uint) (err error) {
	db := DatabaseGet()

	var category model.Category

	count := db.Where("id = ?", categoryId).Find(&category)

	if count.Error != nil {
		zap.L().Error("query post failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query post failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if dir == 1 {
		category.Follow++
	} else {
		category.Follow--
	}

	if err = db.Save(&category).Error; err != nil {
		db.Rollback()
		zap.L().Error("update post speak count failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

// 更新版块发言次数
func CategoryCountSpeak(categoryId uint) (err error) {
	db := DatabaseGet()

	var category model.Category

	count := db.Where("id = ?", categoryId).Find(&category)

	if count.Error != nil {
		zap.L().Error("query post failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query post failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if err = db.Model(&category).Update("speak", category.Speak+1).Error; err != nil {
		db.Rollback()
		zap.L().Error("update post speak count failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

// 更新帖子发言次数
func PostCountSpeak(postId uint) (err error) {
	db := DatabaseGet()

	var post model.Post

	count := db.Where("id = ?", postId).Find(&post)

	if count.Error != nil {
		zap.L().Error("query post failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query post failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if err = db.Model(&post).Update("speak", post.Speak+1).Error; err != nil {
		db.Rollback()
		zap.L().Error("update post speak count failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

// 更新用户登陆ip
func UserUpdateLoginIpv4(userId uint, ip string) (err error) {
	db := DatabaseGet()

	var user model.User

	count := db.Where("id = ?", userId).Find(&user)

	if count.Error != nil {
		zap.L().Error("query user failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query user failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if err = db.Model(&user).Update("last_login_ipv4", ip).Error; err != nil {
		db.Rollback()
		zap.L().Error("update user login ip failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return

}

// 更新用户发言次数
func UserCountSpeak(userId uint) (err error) {
	db := DatabaseGet()

	var user model.User

	count := db.Where("id = ?", userId).Find(&user)

	if count.Error != nil {
		zap.L().Error("query user failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query user failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if err = db.Model(&user).Update("speak", user.Speak+1).Error; err != nil {
		db.Rollback()
		zap.L().Error("update user login ip failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

// 更新用户在版块中的积分
func UserCategoryCount(userId uint, categoryId uint) (err error) {
	db := DatabaseGet()

	var userFollow model.UserFollow

	count := db.Where("user_id = ? AND category_id = ?", userId, categoryId).Find(&userFollow)

	if count.Error != nil {
		zap.L().Error("query user failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		return
	}

	if err = db.Model(&userFollow).Update("count", userFollow.Count+1).Error; err != nil {
		db.Rollback()
		zap.L().Error("update user category count failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}
