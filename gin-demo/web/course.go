package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create course",
	})
}

func CreateCourseV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create course v2",
	})
}

func GetCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get course",
	})
}

func EditCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "edit course",
	})
}

func DeleteCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete course",
	})
}
