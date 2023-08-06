package api

import (
	"exam-online-back-end/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type LoginVo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.RegisteredClaims
	UserID   int    `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
}

func Login(ctx *global.Context) {
	var loginVo LoginVo
	_ = ctx.ShouldBindJSON(&loginVo)

	var admin Admin
	if err := global.DB.Table("admin").Select("id,super").Where("username = ? and password = ?", loginVo.Username, loginVo.Password).Take(&admin).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password error",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID:   admin.Id,
		Username: loginVo.Username,
	})
	tokenString, _ := token.SignedString([]byte("abc"))

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"super": admin.Super,
	})
}
