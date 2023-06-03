package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(100;not null;unique"`
	Password  string `gorm:"size:255;"`
}

func main() {
	db := InitDb()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		//判断账号格式
		if len(name) == 0 {
			name = RandomSring(10)
			//ctx.JSON(422, gin.H{"code": 422, "msg": "手机号码格式错误"})
		}
		//判断手机号格式
		if len(telephone) != 11 {
			ctx.JSON(422, gin.H{"code": 422, "msg": "手机号码必须为11位"})
			return
		}
		//判断密码格式
		if len(password) < 6 {
			ctx.JSON(422, gin.H{"code": 422, "msg": "密码输入必须大于6位"})
			return
		}

		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		log.Println(name, telephone, password)

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run())
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true

	}
	return false
}
func RandomSring(n int) string {
	letters := []byte("aewifdjpoaekjfvpijam210139120939")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
func InitDb() *gorm.DB {
	dariveName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "root"
	charset := "utf8"
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
	return db
}
