package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadEnv() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
    viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error config file : %w", err))
	}

    ApplicationPort = viper.GetString("app.port")
    AllowOrigins = viper.GetStringSlice("gin.cors.allow-origins")
    AllowMethods = viper.GetStringSlice("gin.cors.allow-methods")
    AllowHeaders = viper.GetStringSlice("gin.cors.allow-headers")
}
