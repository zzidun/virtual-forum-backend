package util

import (
	"github.com/spf13/viper"
)

func Config_Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("load config fail")
	}
}
