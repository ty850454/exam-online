package main

import (
	"embed"
	_ "exam-online-back-end/api"
	"exam-online-back-end/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/fs"
	"net/http"
)

func main() {
	router := initServer()
	initDb()
	initRouter()
	fmt.Println("http://127.0.0.1:8080/")
	_ = router.Run(":8080")
}

func initDb() {
	// if utils.FileIsNotExist("./db") {
	// 	err := os.WriteFile("./db", initDatabase, 0666)
	// 	if err != nil {
	// 		panic("failed to init database")
	// 	}
	// }
	var err error
	global.DB, err = gorm.Open(sqlite.Open("file:db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

}

func initServer() *gin.Engine {
	if global.DebugMode {
		global.DebugMode = true
		fmt.Println("运行在debug模式")
	} else {
		fmt.Println("运行在release模式")
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()
	router := gin.Default()

	sub, _ := fs.Sub(dist, "dist")
	staticServer = http.FileServer(http.FS(sub))
	router.NoRoute(staticHandle)

	global.Router = router.Group("/api")
	global.Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}

//go:embed dist
var dist embed.FS
var staticServer http.Handler

func staticHandle(ctx *gin.Context) {
	staticServer.ServeHTTP(ctx.Writer, ctx.Request)
}
