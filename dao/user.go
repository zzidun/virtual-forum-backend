package dao

import (
	"zzidun.tech/vforum0/model"
)

func CheckUserExist(name string, email string) (error error) {
	// sqlstr := `select count(user_id) from user where username = ?`
	// var count int
	// if err := gDatebase.Get(&count, sqlstr, username); err != nil {
	// 	return err
	// }
	// if count > 0 {
	// 	return errors.New("用户已存在")
	// }
	return
}

func UserRegister(user *model.User) (err error) {
	// var count uint
	// db := DatabaseGet()

	// err := CheckUserExist(user.Name, user.Email)
	// if err != nil {
	// 	// 数据库查询出错
	// 	return err
	// }

	// // 2、生成UID
	// userId, err := util.SnowflakeGet()
	// if err != nil {
	// 	return mysql.ErrorGenIDFailed
	// }
	// // 构造一个User实例
	// u := model.User{}
	// // 3、保存进数据库
	// return mysql.InsertUser(u)

	return
}
