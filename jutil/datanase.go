package jutil

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

//获取实例
func NewDb() *gorm.DB {
	var instance *gorm.DB
    var once sync.Once
	once.Do(func() {
		conf:=NewConfig()
		hostname:=conf.GetString("database.d_jcrawl.hostname")
		port:=conf.GetInt("database.d_jcrawl.port")
		username:=conf.GetString("database.d_jcrawl.username")
		password:=conf.GetString("database.d_jcrawl.password")
		database:=conf.GetString("database.d_jcrawl.database")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err!=nil{
			panic("failed to connect database")
		}
		instance=db
	})
	return instance
}
