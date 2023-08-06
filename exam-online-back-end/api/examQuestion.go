package api

import (
	"exam-online-back-end/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ExamQuestion struct {
	Id      int    `json:"id,omitempty"`
	AdminId int    `json:"adminId,omitempty"`
	PaperId int    `json:"paperId,omitempty"`
	GroupId int    `json:"groupId,omitempty"`
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

func CreateExamQuestion(ctx *global.Context) {
	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	paperId := ctx.Param("id")
	paperIdInt, _ := strconv.Atoi(paperId)
	groupId := ctx.Param("groupId")
	groupIdInt, _ := strconv.Atoi(groupId)

	var question ExamQuestion
	_ = ctx.ShouldBindJSON(&question)
	question.PaperId = paperIdInt
	question.PaperId = groupIdInt
	question.AdminId = userId

	if err := global.DB.Save(&question).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": question.Id,
	})
}

func ListExamQuestion(ctx *global.Context) {
	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	ctx.Param("id")
	print(userId)
}
