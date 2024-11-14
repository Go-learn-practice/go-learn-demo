package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	// 初始化课程路由
	initApi(r)
	// 初始化 API 路由
	initCourse(r)
}
