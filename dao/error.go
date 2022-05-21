package dao

import "errors"

var (
	ErrorPasswordWrong  = errors.New("密码错误")
	ErrorInvalidID      = errors.New("无效的ID")
	ErrorQueryFailed    = errors.New("查询数据失败")
	ErrorInsertFailed   = errors.New("插入数据失败")
	ErrorUpdateFailed   = errors.New("更新数据失败")
	ErrorDeleteFailed   = errors.New("删除数据失败")
	ErrorExistFailed    = errors.New("数据已存在")
	ErrorNotExistFailed = errors.New("数据不存在")
)
