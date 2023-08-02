package api

import (
	"errors"
	"exam-online-back-end/global"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
)

type Admin struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Super    bool   `json:"super,omitempty"`
}

func CreateAdmin(ctx *gin.Context) {

	if _, ok := GetUserId(ctx); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	var admin Admin
	_ = ctx.ShouldBindJSON(&admin)
	admin.Super = false
	if err := global.DB.Save(&admin).Error; err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "The same username already exists",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add admin",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": admin.Id,
	})
}

func PageAdmin(ctx *gin.Context) {
	_, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	var result []Admin
	var count int64
	query := global.DB.Model(&result).Where("super = 0")

	if err := query.Count(&count).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if count == 0 {
		ctx.JSON(http.StatusOK, EmptyPage)
		return
	}

	if err := query.Limit(500).Offset(0).Find(&result).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, PageRes{
		Total: count,
		Data:  result,
	})

}

func DeleteAdmin(ctx *gin.Context) {

	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	var admin Admin
	if err := global.DB.First(&admin, userId).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	if !admin.Super {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "Only Super admin can delete other admin.",
		})
		return
	}

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	if userId == id {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "You cannot delete your account.",
		})
		return
	}

	global.DB.Exec("delete from admin where id = ?", id)
	ctx.JSON(http.StatusOK, gin.H{})
}
