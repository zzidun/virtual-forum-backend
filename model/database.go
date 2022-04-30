package model

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var g_db *gorm.DB

func DB_Init() {

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

	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&AdminGroup{})
	db.AutoMigrate(&AdminGroupPerm{})
	db.AutoMigrate(&AdminGroupMember{})
	db.AutoMigrate(&AdminOperator{})

	db.AutoMigrate(&Category{})
	db.AutoMigrate(&CategoryAdmin{})
	db.AutoMigrate(&CategoryUser{})

	db.AutoMigrate(&Post{})

	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&CommentInfo{})

	db.AutoMigrate(&UserInfo{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserOperator{})
	db.AutoMigrate(&UserBanned{})
	db.AutoMigrate(&UserShield{})
	db.AutoMigrate(&UserCollect{})
	db.AutoMigrate(&UserFollow{})

	db.AutoMigrate(&BannedIpv4{})
	db.AutoMigrate(&FailedUser{})
	db.AutoMigrate(&Report{})

	db.AutoMigrate(&Message{})

	g_db = db
}

func DB_Get() *gorm.DB {
	return g_db
}
