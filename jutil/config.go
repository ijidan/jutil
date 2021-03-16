package jutil

import (
	"github.com/spf13/viper"
	"os"
	"sync"
)


//获取配置
func NewConfig() *viper.Viper {
	var once sync.Once
	var instance *viper.Viper
	once.Do(func() {
		wd, _ := os.Getwd()
		pathSep := GetPathSep()
		confPath := wd + pathSep + "config"
		instance = viper.New()
		instance.AddConfigPath(confPath)
		instance.SetConfigName("default")
		instance.SetConfigType("json")
		if err := instance.ReadInConfig(); err != nil {
			panic(err)
		}
	})
	return instance
}

//获取文件分隔符
func GetPathSep() string {
	pathSep := string(os.PathSeparator)
	return pathSep
}
