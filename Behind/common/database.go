package common

import (
	"Behind/mods"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

// 初始化数据库
func InitDb(DariveName string, Host string, Port string, Database string, Username string, Password string, Charset string) *gorm.DB {
	log.Println(DariveName, Host, Port, Database, Username, Password, Charset)
	dariveName := DariveName
	host := Host
	port := Port
	database := Database
	username := Username
	password := Password
	charset := Charset
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(dariveName, args)

	if err != nil {
		panic("错误:" + err.Error())

	}
	db.AutoMigrate(&mods.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
