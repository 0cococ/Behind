package controller

import (
	"Behind/common"
	"Behind/mods"
	"Behind/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
)

//func Register(ctx *gin.Context) {
//	DB := common.GetDB()
//	//获取参数
//	name := ctx.PostForm("name")
//	telephone := ctx.PostForm("telephone")
//	password := ctx.PostForm("password")
//	log.Println("1", telephone)
//	//判断账号格式
//	if len(name) == 0 {
//		name = util.RandomSring(10)
//		//ctx.JSON(422, gin.H{"code": 422, "msg": "手机号码格式错误"})
//	}
//
//	//判断手机号格式
//	if len(telephone) != 11 {
//		ctx.JSON(422, gin.H{"code": 422, "msg": "手机号码必须为11位"})
//		return
//	}
//	//判断密码格式
//	if len(password) < 6 {
//		ctx.JSON(422, gin.H{"code": 422, "msg": "密码输入必须大于6位"})
//		return
//	}
//
//	if isTelephoneExist(DB, telephone) {
//		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
//		return
//	}
//	//创建用户
//	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "创建用户失败,请重新注册"})
//		return
//	}
//	newUser := mods.User{
//		Name:      name,
//		Telephone: telephone,
//		Password:  string(hasedPassword),
//	}
//	DB.Create(&newUser)
//
//	log.Println(name, telephone, password)
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "注册成功",
//	})
//}
//	func Login(ctx *gin.Context) {
//		DB := common.GetDB()
//		telephone := ctx.PostForm("telephone")
//		password := ctx.PostForm("password")
//		if len(telephone) != 11 {
//			ctx.JSON(422, gin.H{"code": 422, "msg": "手机号码必须为11位"})
//			return
//		}
//
//		if len(password) < 6 {
//			ctx.JSON(422, gin.H{"code": 422, "msg": "密码输入必须大于6位"})
//			return
//		}
//		//判断手机号是否注册
//		var user mods.User
//		DB.Where("telephone=?", telephone).First(&user)
//		if user.ID == 0 {
//			ctx.JSON(422, gin.H{"code": 422, "msg": "用户不存在"})
//			return
//		}
//		//判断密码是否正确
//		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
//			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
//			return
//		}
//
//		token, err := common.ReleaseToken(user)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统错误"})
//			return
//		}
//
//		ctx.JSON(
//			200,
//			gin.H{
//				"code": 200, "data": gin.H{"token": token}, "msg": "登录成功"})
//
// }
func Register(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	log.Println("1", telephone)

	var errors []string

	// 判断账号格式
	if len(name) == 0 {
		name = util.RandomSring(10)
	}

	// 判断手机号格式
	if len(telephone) != 11 {
		errors = append(errors, "手机号码必须为11位")
	}

	// 判断密码格式
	if len(password) < 6 {
		errors = append(errors, "密码输入必须大于6位")
	}

	// 检查手机号是否已存在
	if isTelephoneExist(DB, telephone) {
		errors = append(errors, "用户已经存在")
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errors = append(errors, "创建用户失败，请重新注册")
	}

	// 检查是否有错误消息
	if len(errors) > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "errors": errors})
		return
	}

	newUser := mods.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}

	// 使用数据库事务创建用户
	err = DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newUser).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "创建用户失败，请重新注册"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 错误消息切片
	var errors []string

	// 验证手机号码
	if matched, _ := regexp.MatchString(`^\d{11}$`, telephone); !matched {
		errors = append(errors, "手机号码格式不正确")
	}

	// 验证密码长度
	if len(password) < 6 {
		errors = append(errors, "密码输入必须大于6位")
	}

	// 如果存在错误则返回
	if len(errors) > 0 {
		ctx.JSON(422, gin.H{"code": 422, "errors": errors})
		return
	}

	// 检查用户是否存在于数据库中
	var user mods.User
	if err := DB.Where("telephone = ?", telephone).First(&user).Error; err != nil {
		ctx.JSON(422, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	// 比较密码的哈希值
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 生成令牌
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统错误"})
		return
	}

	// 返回成功响应
	ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "登录成功"})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user mods.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true

	}
	return false
}
