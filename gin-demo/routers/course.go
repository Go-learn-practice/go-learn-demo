package routers

import (
	"gin-demo/middleware"
	"gin-demo/web"
	"github.com/gin-gonic/gin"
)

func initCourse(r *gin.Engine) {
	// http://localhost:8080/v1/course
	v1 := r.Group("/v1", middleware.TokenCheck, middleware.AuthCheck)
	v1.POST("/course", web.CreateCourse)
	v1.GET("/course", web.GetCourse)
	v1.PUT("/course", web.EditCourse)
	v1.DELETE("/course", web.DeleteCourse)

	v2 := r.Group("/v2")
	v2.POST("/course", web.CreateCourseV2)
}
