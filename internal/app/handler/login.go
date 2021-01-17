package handler

import (
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {

}

func generateToken(c *gin.Context, user model.User) (string, error) {
	j := model.NewJWT()

	claims := model.CustomClaims{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + 3600,
			Issuer:    "chaosi-zju",
		},
	}

	return j.CreateToken(claims)
}
