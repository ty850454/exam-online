package api

import (
	"exam-online-back-end/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExamQuestion struct {
	Id      int    `json:"id,omitempty"`
	AdminId int    `json:"adminId,omitempty"`
	PaperId int    `json:"paperId,omitempty"`
	Title   string `json:"title,omitempty"`
	Type    int    `json:"type,omitempty"`
	Answer  int    `json:"answer,omitempty"`
}

type ExamOption struct {
	Id         int `json:"id,omitempty"`
	AdminId    int `json:"adminId,omitempty"`
	PaperId    int `json:"paperId,omitempty"`
	QuestionId int `json:"questionId,omitempty"`
}

func CreateExamQuestion(ctx *gin.Context) {
	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	var examQuestion ExamQuestion
	_ = ctx.ShouldBindJSON(&examQuestion)
	examQuestion.AdminId = userId

	if err := global.DB.Save(&examQuestion).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": examQuestion.Id,
	})
}

func ListExamQuestion(ctx *gin.Context) {
	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	ctx.Param("id")
	print(userId)
}
