package api

import (
	"exam-online-back-end/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ExamPaper struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Type  int    `json:"type,omitempty"`
	// 1=未发布，2=已发布
	JoinType int `json:"joinType,omitempty"`
	AdminId  int `json:"adminId,omitempty"`
	// 1=未发布，2=已发布
	Status    int       `json:"status,omitempty"`
	StartTime *DateTime `json:"startTime,omitempty"`
	EndTime   *DateTime `json:"endTime,omitempty"`
}

func CreateExamPaper(ctx *gin.Context) {
	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	var exam ExamPaper
	_ = ctx.ShouldBindJSON(&exam)
	exam.Status = 1
	exam.AdminId = userId
	if err := global.DB.Save(&exam).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add exam",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": exam.Id,
	})
}

type ExamPaperReq struct {
	PageReq
}

func UserGetExamPaper(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	var result ExamPaper
	result.Id = idInt
	if err := global.DB.Find(&result).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	result.AdminId = 0
	result.JoinType = 0
	ctx.JSON(http.StatusOK, result)
}

func ListExamPaper(ctx *gin.Context) {
	userId, ok := GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
		})
		return
	}

	var param ExamPaperReq
	err := ctx.BindQuery(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var result []ExamPaper
	var count int64
	query := global.DB.Model(&result).Where("admin_id = ?", userId)

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

	if err := query.Limit(param.PageSize).Offset(param.GetOffset()).Find(&result).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i := range result {
		result[i].AdminId = 0
	}

	ctx.JSON(http.StatusOK, PageRes{
		Total: count,
		Data:  result,
	})
	return
}
