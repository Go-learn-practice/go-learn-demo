package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginReq struct {
	Username string `form:"user_name" json:"username"`
	Pwd      string `form:"pwd" json:"pwd"`
}

func Login(c *gin.Context) {
	req := loginReq{}
	c.Bind(&req)

	c.JSON(http.StatusOK, req)
}
