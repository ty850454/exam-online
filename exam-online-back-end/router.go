package main

import (
	"exam-online-back-end/api"
	"exam-online-back-end/global"
)

func initRouter() {
	global.ApiRouter.POST("/login", api.Login)

	global.ApiRouter.GET("/admin", api.PageAdmin)
	global.ApiRouter.POST("/admin", api.CreateAdmin)
	global.ApiRouter.DELETE("/admin/:id", api.DeleteAdmin)

	global.ApiRouter.POST("/exam/:paperId", api.SaveExam)
	global.ApiRouter.GET("/exam/:paperId", api.GetExam)

	// global.ApiRouter.POST("/examPaper", api.CreateExamPaper)
	// global.ApiRouter.GET("/examPaper", api.ListExamPaper)

	// global.ApiRouter.POST("/exam/:id/group", api.CreateExamPaperGroup)
	// global.ApiRouter.DELETE("/exam/:id/group/:groupId", api.DeleteExamPaperGroup)
	// global.ApiRouter.PUT("/exam/:id/group/:groupId", api.UpdateExamPaperGroup)
	// global.ApiRouter.GET("/exam/:id/group", api.ListExamPaperGroup)
	//
	// global.ApiRouter.POST("/exam/:id/group/:groupId/question", api.CreateExamQuestion)
	// global.ApiRouter.DELETE("/exam/:id/group/:groupId/question/:questionId", api.DeleteExamQuestion)
	// global.ApiRouter.PUT("/exam/:id/group/:groupId/question/:questionId", api.UpdateExamQuestion)
	// global.ApiRouter.GET("/exam/:id/group/question", api.ListExamQuestion)

	global.UserRouter.GET("/exam/:id", api.UserGetExamPaper)

	global.ApiRouter.POST("/config", api.AddConfig)
	global.ApiRouter.GET("/config/:group", api.ListConfig)
	global.ApiRouter.GET("/config/:group/:name", api.GetConfig)
	global.ApiRouter.PATCH("/config/:id", api.UpdateConfig)
	global.ApiRouter.DELETE("/config/:id", api.DeleteConfig)

}
