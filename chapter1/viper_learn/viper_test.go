package main

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestGetConfig(t *testing.T) {

	viper.SetDefault("AppName", "default_app") //设置默认值
	viper.SetDefault("Env", "local")

	fmt.Println(viper.GetString("AppName")) //获取值
	fmt.Println(viper.GetString("Env"))

	viper.SetConfigName("config")     //设置配置名称
	viper.SetConfigType("yaml")       //设置文件类型
	viper.AddConfigPath("../config/") //设置文件路径

	err := viper.ReadInConfig() //读取文件
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(viper.GetString("AppName")) //获取数据
	fmt.Println(viper.GetString("Env"))
	fmt.Println(viper.GetString("Computer"))
}
