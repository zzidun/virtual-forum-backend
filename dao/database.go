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

func DatabaseInit() {

	// driverName := viper.Get("database.driver_name")

	host := viper.Get("data_source.host")
	port := viper.Get("data_source.port")
	database := viper.Get("data_source.database")
	username := viper.Get("data_source.username")
	password := viper.Get("data_source.password")
	charset := viper.Get("data_source.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.AdminGroup{})

	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.CategoryAdmin{})

	db.AutoMigrate(&model.Post{})

	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.CommentInfo{})

	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.BannedIpv4{})
	db.AutoMigrate(&model.FailedUser{})
	db.AutoMigrate(&model.UserFollow{})
	db.AutoMigrate(&model.UserShield{})
	db.AutoMigrate(&model.UserCollect{})

	gDatebase = db
}

func DatabaseGet() *gorm.DB {
	return gDatebase
}
