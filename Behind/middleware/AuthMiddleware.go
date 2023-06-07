package middleware

import (
	"Behind/common"
	"Behind/mods"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		fmt.Println(tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"CODE": 401, "msg": "权限不足1"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"CODE": 401, "msg": "权限不足2"})
			ctx.Abort()
			return
		}

		userId := claims.UserId
		DB := common.GetDB()
		var user mods.User
		DB.First(&user, userId)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"CODE": 401, "msg": "权限不足3"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()

	}

}
