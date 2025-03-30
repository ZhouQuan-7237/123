package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	fmt.Println("服务器已启动，等待客户端连接...")
	// router.Use(middlewares...) // 使用中间件
	return router
}
