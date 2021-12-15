package util

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zzidun.tech/vforum0/model"
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

	db.AutoMigrate(&model.User{})

	g_db = db
}

func DB_Get() *gorm.DB {
	return g_db
}
