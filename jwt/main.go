package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"jwt/jwtV1"
	"time"
)

/*
1. 选择一种加密方式
2. 构造payload部分
3. 执行加密过程
4. 返回 jwt 令牌
*/

type MyClaims struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	jwt.RegisteredClaims
}

func main() {
	myClaims := MyClaims{
		Name:   "nick",
		Gender: "male",
		Age:    20,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	var outClaims MyClaims
	var sign string
	var err error

	hs := jwtV1.HS{
		Key: "your-256-bit-secret",
	}
	sign, err = hs.Encode(myClaims)
	fmt.Println(sign, err)

	err = hs.Decode(sign, &outClaims)
	fmt.Println(err, outClaims)
}
