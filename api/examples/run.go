package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-admin/common/global"
	mycasbin "go-admin/pkg/casbin"
	"go-admin/pkg/logger"
)

func main() {
	var err error
	global.Eloquent, err = gorm.Open(mysql.Open("root:123456@tcp/inmg?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mycasbin.Setup()
	logger.Setup()
	global.GinEngine = gin.Default()
	global.GinEngine.Use(MiddleWareA())
	//router.InitRouter()
	log.Fatal(global.GinEngine.Run(":8000"))
}

func MiddleWareA() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
