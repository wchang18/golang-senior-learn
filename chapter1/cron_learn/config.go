package cron_learn

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName   string `mapstructure:"AppName"`
	Env       string `mapstructure:"Env"`
	Scheduler []Task `mapstructure:"Scheduler"`
}

var config Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(&config)
}
