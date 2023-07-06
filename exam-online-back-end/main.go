package main

import (
	"embed"
	"exam-online-back-end/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

func main() {
	initServer()
	fmt.Println("http://127.0.0.1:8080/")
	_ = global.Router.Run(":8080")
}

func initServer() {
	if global.DebugMode {
		global.DebugMode = true
		fmt.Println("运行在debug模式")
	} else {
		fmt.Println("运行在release模式")
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()
	global.Router = gin.Default()

	sub, _ := fs.Sub(dist, "dist")
	staticServer = http.FileServer(http.FS(sub))

	global.Router.NoRoute(staticHandle)

	global.Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

//go:embed dist
var dist embed.FS
var staticServer http.Handler

func staticHandle(ctx *gin.Context) {
	staticServer.ServeHTTP(ctx.Writer, ctx.Request)
}
