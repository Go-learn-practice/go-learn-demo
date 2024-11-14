package routers

import (
	"gin-demo/web"
	"github.com/gin-gonic/gin"
)

func initApi(r *gin.Engine) {
	// http://localhost:8080/api/v1
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/ping", web.Ping)
	v1.POST("/login", web.Login)
	v1.POST("/register", web.Register)
}
