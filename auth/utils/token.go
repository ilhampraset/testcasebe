package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
	ExtractToken(tokenBearer string) (map[string]interface{}, error)
}

type authCustomClaims struct {
	Username string `json:"username"`
	User     bool   `json:"user"`

	jwt.StandardClaims
}

type JwtUtil struct {
	secretKey string
	issure    string
}

func JWTAuth() Token {
	return &JwtUtil{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getSecretKey() string {

	return "secret"
}

func (j *JwtUtil) GenerateToken(username string, isUser bool) string {
	claims := &authCustomClaims{
		username,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    j.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *JwtUtil) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})

}

func (j *JwtUtil) ExtractToken(tokenBearer string) (map[string]interface{}, error) {
	tokenString := strings.SplitAfter(tokenBearer, "Bearer")[1]
	token, _ := j.ValidateToken(strings.TrimSpace(tokenString))
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
