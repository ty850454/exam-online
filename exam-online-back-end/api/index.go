package api

import (
	"exam-online-back-end/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
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

func Login(ctx *gin.Context) {
	var loginVo LoginVo
	_ = ctx.ShouldBindJSON(&loginVo)

	var admin Admin
	if err := global.DB.Table("admin").Select("id").Where("username = ? and password = ?", loginVo.Username, loginVo.Password).Take(&admin).Error; err != nil {
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
	})
}

func GetUserId(ctx *gin.Context) (int, bool) {
	authorization := ctx.Request.Header.Get("Authorization")
	if len(authorization) == 0 {
		return 0, false
	}
	index := strings.IndexByte(authorization, ' ')
	if index < 0 {
		return 0, false
	}
	_type := authorization[:index]
	if strings.Compare(_type, "Bearer") != 0 {
		return 0, false
	}
	index = strings.LastIndexByte(authorization, ' ') + 1
	tokenStr := authorization[index:]

	token, _ := jwt.Parse(tokenStr, checkToken)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims["userId"].(float64); ok {
			return int(userId), true
		}
	}
	return 0, false
}

func checkToken(_ *jwt.Token) (interface{}, error) {
	return []byte("abc"), nil
}
