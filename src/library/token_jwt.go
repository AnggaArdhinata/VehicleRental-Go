package library

import (
	
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var myKey = []byte(os.Getenv("JWT_KEY"))

type claims struct {
	User_Id uint
	Role string
	jwt.StandardClaims
}

func NewToken(id uint, role string) *claims {
	return &claims{
		User_Id: id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}
}

func (c *claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	
	return tokens.SignedString(myKey)
}

func CheckToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims := tokens.Claims.(*claims)
	return claims, nil
}