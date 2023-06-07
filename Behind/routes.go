package main

import (
	"Behind/controller"
	"Behind/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) gin.IRoutes {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
