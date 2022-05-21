package dao

import "zzidun.tech/vforum0/model"

func BannedIpv4Create(left string, right string, userId uint) (ban *model.BannedIpv4, err error) {
	ban = &model.BannedIpv4{
		LeftIp:  left,
		RightIp: right,
		UserId:  userId,
	}

	db := DatabaseGet()
	if db.Create(&ban).Error != nil {
		err = ErrorInsertFailed
		return
	}

	return
}

func BannedIpv4Delete(banId uint) (err error) {

	db := DatabaseGet()

	ban, err := BannedIpv4QueryById(banId)
	if err != nil {
		return
	}

	if err = db.Delete(&ban).Error; err != nil {
		err = ErrorDeleteFailed
		return
	}

	return
}

func BannedIpv4Update(banId uint, left string, right string, userId uint) (err error) {
	db := DatabaseGet()

	ban, err := BannedIpv4QueryById(banId)
	if err != nil {
		return
	}

	ban.LeftIp = left
	ban.RightIp = right
	ban.UserId = userId

	if db.Save(&ban).Error != nil {
		return ErrorUpdateFailed
	}

	return
}

func BannedIpv4Query(left int, right int) (banList []*model.BannedIpv4, totNum int64, curNum int64, err error) {
	db := DatabaseGet()

	count := db.Limit(right - left).Offset(left).Find(&banList)

	if count.Error != nil {
		err = ErrorQueryFailed
	}
	curNum = count.RowsAffected
	db.Model(&model.BannedIpv4{}).Count(&totNum)

	return
}

func BannedIpv4QueryById(banId uint) (ban *model.BannedIpv4, err error) {

	db := DatabaseGet()

	count := db.Where("id = ?", banId).Find(&banId)

	if count.Error != nil {
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		err = ErrorNotExistFailed
		return
	}

	return
}

func FailedCount() {

}

func FailedQuery() {

}
