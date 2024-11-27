package jwtV1

import "github.com/golang-jwt/jwt/v5"

type ED struct {
	PrivateKey string
	PublicKey  string
}

func (ed *ED) Encode(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	pKey, err := jwt.ParseEdPrivateKeyFromPEM([]byte(ed.PrivateKey))
	if err != nil {
		return "", err
	}
	sign, err := token.SignedString(pKey)
	if err != nil {
		return "", err
	}
	return sign, nil
}

func (ed *ED) Decode(sign string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseEdPublicKeyFromPEM([]byte(ed.PublicKey))
	})
	return err
}
