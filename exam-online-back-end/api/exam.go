package api

import (
	"encoding/json"
	"exam-online-back-end/global"
)

// type ExamVo struct {
// 	Id       int
// 	Title    string
// 	Type     int
// 	JoinType int
// 	Status   int
// 	Groups   []ExamGroupVo
// }

type ExamGroupVo struct {
	Title          string           `json:"title,omitempty"`
	QuestionNumber int              `json:"questionNumber"`
	Score          int              `json:"score"`
	Questions      []ExamQuestionVo `json:"questions,omitempty"`
}

type ExamQuestionVo struct {
	Title   string         `json:"title,omitempty"`
	Type    int            `json:"type,omitempty"`
	Answer  int            `json:"answer,omitempty"`
	Score   int            `json:"score"`
	Options []ExamOptionVo `json:"options,omitempty"`
}

type ExamOptionVo struct {
	Code     string `json:"code,omitempty"`
	Title    string `json:"title,omitempty"`
	IsAnswer bool   `json:"isAnswer,omitempty"`
}

var codes = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func SaveExam(ctx *global.Context) {
	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.ErrorInvalidToken()
		return
	}

	paperId := ctx.ParamInt("paperId")
	var groups []ExamGroupVo
	_ = ctx.ShouldBindJSON(&groups)

	for i, group := range groups {
		questionNumber := len(group.Questions)
		groups[i].QuestionNumber = questionNumber
		score := 0
		for qi, question := range group.Questions {
			score += question.Score

			for oi := range question.Options {
				groups[i].Questions[qi].Options[oi].Code = codes[oi]
			}
		}
		groups[i].Score = score
	}

	j, _ := json.Marshal(groups)
	question := string(j)

	if err := global.DB.Model(&ExamPaper{}).Where("id = ? and admin_id = ?", paperId, userId).Update("question", question).Error; err != nil {
		ctx.ErrorBad(err.Error())
		return
	}

	// _ = json.Unmarshal(j, &groups)
	//
	// ctx.Ok(groups)
}

func GetExam(ctx *global.Context) {
	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.ErrorInvalidToken()
		return
	}

	paperId := ctx.ParamInt("paperId")

	var examPaper ExamPaper
	if err := global.DB.Where("id = ? and admin_id = ?", paperId, userId).Find(&examPaper).Error; err != nil {
		ctx.ErrorBad(err.Error())
		return
	}

	var groups []ExamGroupVo
	if len(examPaper.Question) > 0 {
		_ = json.Unmarshal([]byte(examPaper.Question), &groups)
	} else {
		groups = make([]ExamGroupVo, 0)
	}

	ctx.Ok(groups)
}
