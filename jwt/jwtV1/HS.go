package jwtV1

import "github.com/golang-jwt/jwt/v5"

type HS struct {
	Key        string
	SignMethod HSSignMethod
}

type HSSignMethod string

const (
	HS256 HSSignMethod = "HS256"
	HS384 HSSignMethod = "HS384"
	HS512 HSSignMethod = "HS512"
)

func (hs *HS) getHSSignMethod(name HSSignMethod) *jwt.SigningMethodHMAC {
	switch name {
	case HS256:
		return jwt.SigningMethodHS256
	case HS384:
		return jwt.SigningMethodHS384
	case HS512:
		return jwt.SigningMethodHS512
	}
	return jwt.SigningMethodHS256
}

func (hs *HS) Encode(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(hs.getHSSignMethod(hs.SignMethod), claims)
	sign, err := token.SignedString([]byte(hs.Key))
	if err != nil {
		return "", err
	}
	return sign, nil
}

func (hs *HS) Decode(sign string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(hs.Key), nil
	})
	return err
}
