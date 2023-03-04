package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadEnv() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error config file : %w", err))
	}
}
