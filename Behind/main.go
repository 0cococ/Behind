package main

import (
	"Behind/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDb("mysql", "192.168.86.131", "3306", "ginessential", "fjc", "123456", "utf8")
	defer db.Close()
	r := gin.Default()
	CollectRoute(r)
	panic(r.Run())
}
