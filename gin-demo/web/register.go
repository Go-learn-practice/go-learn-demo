package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerReq struct {
	UserName string `form:"username" binding:"required"`
	Pwd      string `form:"pwd" binding:"required"`
	Phone    string `form:"phone" binding:"required,e164"`
	Email    string `form:"email" binding:"omitempty,email"`
}

func Register(c *gin.Context) {
	req := &registerReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
}
