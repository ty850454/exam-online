package api

import (
	"exam-online-back-end/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ExamGroup struct {
	Id             int    `json:"id,omitempty"`
	PaperId        int    `json:"paperId,omitempty"`
	Title          string `json:"title,omitempty"`
	AdminId        int    `json:"adminId,omitempty"`
	QuestionNumber int    `json:"questionNumber"`
	Score          int    `json:"score"`
}

func CreateExamPaperGroup(ctx *global.Context) {
	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	paperId := ctx.Param("id")
	paperIdInt, _ := strconv.Atoi(paperId)

	var group ExamGroup
	_ = ctx.ShouldBindJSON(&group)
	group.PaperId = paperIdInt
	group.AdminId = userId
	group.QuestionNumber = 0
	group.Score = 0
	if err := global.DB.Save(&group).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add group",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": group.Id,
	})
}

func DeleteExamPaperGroup(ctx *global.Context) {
	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}
	groupId := ctx.Param("groupId")

	if err := global.DB.Where("id = ? and admin_id = ?", groupId, userId).Delete(&ExamGroup{}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func UpdateExamPaperGroup(ctx *global.Context) {
	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	groupId := ctx.Param("groupId")
	groupIdInt, _ := strconv.Atoi(groupId)
	paperId := ctx.Param("id")
	paperIdInt, _ := strconv.Atoi(paperId)

	var group ExamGroup
	_ = ctx.ShouldBindJSON(&group)
	group.Id = groupIdInt
	group.PaperId = paperIdInt
	group.AdminId = userId
	// TODO 查询
	group.QuestionNumber = 0
	group.Score = 0

	if err := global.DB.Save(group).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func ListExamPaperGroup(ctx *global.Context) {
	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	paperId := ctx.Param("id")

	var result []ExamGroup
	if err := global.DB.Where("admin_id = ? and `paper_id` = ?", userId, paperId).Find(&result).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i := range result {
		result[i].AdminId = 0
	}

	ctx.JSON(http.StatusOK, result)
}
