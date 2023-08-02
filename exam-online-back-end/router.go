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

	global.ApiRouter.POST("/exam", api.CreateExamPaper)
	global.ApiRouter.GET("/exam", api.ListExamPaper)

	global.UserRouter.GET("/exam/:id", api.UserGetExamPaper)

	global.ApiRouter.POST("/config", api.AddConfig)
	global.ApiRouter.GET("/config/:group", api.ListConfig)
	global.ApiRouter.GET("/config/:group/:name", api.GetConfig)
	global.ApiRouter.PATCH("/config/:id", api.UpdateConfig)
	global.ApiRouter.DELETE("/config/:id", api.DeleteConfig)

}
