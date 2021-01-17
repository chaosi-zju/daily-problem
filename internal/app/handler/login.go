package handler

import (
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	var param model.LoginParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseError(c, 500, "param invalid")
	} else {
		if user, err := model.LoginCheck(param); err != nil {
			util.ResponseError(c, 500, err.Error())
		} else {
			if token, err := generateToken(user); err != nil {
				util.ResponseError(c, 500, err.Error())
			} else {
				util.ResponseSuccess(c, 200, "success", map[string]string{
					"token": token,
				})
			}
		}
	}
}

func Register(c *gin.Context) {
	var param model.RegisterParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseError(c, 500, "param invalid")
	} else {
		if user, err := model.Register(param); err != nil {
			util.ResponseError(c, 500, "register failed: "+err.Error())
		} else {
			if token, err := generateToken(user); err != nil {
				util.ResponseError(c, 500, err.Error())
			} else {
				util.ResponseSuccess(c, 200, "success", map[string]string{
					"token": token,
				})
			}
		}
	}
}

func generateToken(user model.User) (string, error) {
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
