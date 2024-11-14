package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthCheck 权限检查
func AuthCheck(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userName, _ := c.Get("user_name")
	fmt.Printf("auth check userID: %s, userName: %s\n", userId, userName)
	c.Next()
}

var token = "123456"

func TokenCheck(c *gin.Context) {
	accessToken := c.Request.Header.Get("access-token")
	if accessToken != token {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token 检查失败",
		})
		c.AbortWithError(http.StatusInternalServerError, errors.New("token 检查失败"))
	}
	c.Set("user_name", "nick")
	c.Set("user_id", "10001")
	c.Next()
}
