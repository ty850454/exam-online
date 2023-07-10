package main

import (
	"exam-online-back-end/api"
	"exam-online-back-end/global"
)

func initRouter() {
	global.Router.POST("/admin", api.Create)
}
