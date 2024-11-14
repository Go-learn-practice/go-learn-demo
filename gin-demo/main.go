package main

import (
	"gin-demo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	routers.InitRouters(r)
	r.Run(":8080")
}
