package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ilhampraset/testcasebe/auth/utils"
	"github.com/ilhampraset/testcasebe/domain"
)

type AuthHandler struct {
	uc  domain.UserUseCase
	jwt utils.Token
}

func NewAuthHandler(uc domain.UserUseCase, jwt utils.Token) *AuthHandler {
	h := &AuthHandler{uc, jwt}

	return h
}

func (h *AuthHandler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	verif, _ := h.uc.VerifyLogin(username, password)

	if !verif {
		c.JSON(401, gin.H{"message": "invalid credentials"})
	} else {
		c.JSON(200, gin.H{"access_token": h.jwt.GenerateToken(username, true)})
	}

}

func (h *AuthHandler) Me(c *gin.Context) {
	auth, _ := h.jwt.ExtractToken(c.GetHeader("Authorization"))
	user, _ := h.uc.Me(fmt.Sprintf("%v", auth["username"]))
	c.JSON(200, gin.H{"data": user})

}
