package api

import (
	"exam-online-back-end/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Config struct {
	Id      int    `json:"id,omitempty"`
	Group   string `json:"group,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
	AdminId int    `json:"adminId,omitempty"`
}

func GetConfig(ctx *gin.Context) {
	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	group := ctx.Param("group")
	name := ctx.Param("name")

	var result Config
	if err := global.DB.Where("admin_id = ? and `group` = ? and `name` = ?", userId, group, name).Limit(1).Find(&result).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result.AdminId = 0
	result.Group = ""

	ctx.JSON(http.StatusOK, result)
}

func AddConfig(ctx *gin.Context) {
	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	var config Config
	_ = ctx.ShouldBindJSON(&config)
	config.AdminId = userId
	if err := global.DB.Save(&config).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add config",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": config.Id,
	})
}
func UpdateConfig(ctx *gin.Context) {
	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	id := ctx.Param("id")

	newName := ctx.Query("newName")
	newValue := ctx.Query("newValue")

	updates := map[string]any{}
	if len(newName) > 0 {
		updates["name"] = newName
	}
	if len(newValue) > 0 {
		updates["value"] = newValue
	}

	if len(updates) > 0 {
		if err := global.DB.Table("config").Where("id = ? and admin_id = ?", id, userId).Updates(updates).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	idInt, _ := strconv.Atoi(id)
	ctx.JSON(http.StatusOK, gin.H{
		"id": idInt,
	})
}

func ListConfig(ctx *gin.Context) {
	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	group := ctx.Param("group")

	var result []Config
	if err := global.DB.Where("admin_id = ? and `group` = ?", userId, group).Find(&result).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i := range result {
		result[i].AdminId = 0
		result[i].Group = ""
	}

	ctx.JSON(http.StatusOK, result)
}

func DeleteConfig(ctx *gin.Context) {
	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}
	id := ctx.Param("id")

	if err := global.DB.Where("id = ? and admin_id = ?", id, userId).Delete(&Config{}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
