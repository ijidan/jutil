package jutil

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//获取实例
func NewDb(dbName string) *gorm.DB {
	var databasePrefix = "database." + dbName + "."
	conf := NewConfig()
	hostname := conf.GetString(databasePrefix + "hostname")
	port := conf.GetInt(databasePrefix + "port")
	username := conf.GetString(databasePrefix + "username")
	password := conf.GetString(databasePrefix + "password")
	database := conf.GetString(databasePrefix + "database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

//获取Partner
func NewDbPartner()  *gorm.DB{
	db:=NewDb("d_partner")
	return db
}
