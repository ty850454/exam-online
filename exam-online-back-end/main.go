package main

import (
	"embed"
	"exam-online-back-end/global"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/fs"
	"net/http"
	"time"
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
	router.Use(cors.New(cors.Config{
		// 准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		// 准许使用的请求方式
		AllowMethods: []string{"*"},
		// 准许使用的请求表头
		AllowHeaders: []string{"*"},
		// 显示的请求表头
		ExposeHeaders: []string{"*"},
		// 凭证共享,确定共享
		AllowCredentials: true,
		// 容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		// 超时时间设定
		MaxAge: 24 * time.Hour,
	}))

	sub, _ := fs.Sub(dist, "dist")
	staticServer = http.FileServer(http.FS(sub))
	router.NoRoute(staticHandle)

	global.ApiRouter = router.Group("/api")
	global.UserRouter = router.Group("/u")

	return router
}

//go:embed dist
var dist embed.FS
var staticServer http.Handler

func staticHandle(ctx *gin.Context) {
	staticServer.ServeHTTP(ctx.Writer, ctx.Request)
}
