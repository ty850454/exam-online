package api

import (
	"exam-online-back-end/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Super    bool   `json:"super,omitempty"`
}

func Create(ctx *gin.Context) {
	var admin Admin
	_ = ctx.ShouldBindJSON(&admin)

	global.DB.Save(&admin)
	ctx.JSON(http.StatusOK, gin.H{
		"id": admin.Id,
	})
}
