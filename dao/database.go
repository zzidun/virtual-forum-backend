package dao

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zzidun.tech/vforum0/model"
)

var gDatebase *gorm.DB

// 初始化数据库，建表
func Init() {

	// driverName := viper.Get("database.driver_name")

	host := viper.Get("mysql.host")
	port := viper.Get("mysql.port")
	database := viper.Get("mysql.database")
	username := viper.Get("mysql.username")
	password := viper.Get("mysql.password")
	charset := viper.Get("mysql.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Categoryer{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.BannedIpv4{})
	db.AutoMigrate(&model.FailedUser{})
	db.AutoMigrate(&model.UserFollow{})
	db.AutoMigrate(&model.UserShield{})
	db.AutoMigrate(&model.UserCollect{})

	gDatebase = db

	AdminInit()
}

func DatabaseGet() *gorm.DB {
	return gDatebase
}

func AdminInit() {
	db := DatabaseGet()

	var user *model.User

	count := db.Find(&model.User{})

	name, emailOrigin, passwordOrigin, err := rootConfig()
	if err != nil {
		return
	}

	if count.RowsAffected == 0 {
		user, err = UserCreate(name, emailOrigin, passwordOrigin)
		if err != nil {
			return
		}

		admin, err := AdminCreate(user.ID)
		if err != nil {
			return
		}

		err = AdminUpdate(admin.ID, 1, 1, 1)
		if err != nil {
			return
		}

	}

	return
}

func rootConfig() (name string, email string, password string, err error) {
	name = viper.Get("root.name").(string)
	email = viper.Get("root.email").(string)
	password = viper.Get("root.password").(string)

	return
}
