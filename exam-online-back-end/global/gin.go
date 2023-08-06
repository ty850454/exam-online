package global

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
)

type Engine struct {
	*gin.Engine
}

type Context struct {
	*gin.Context
}

type RouterGroup struct {
	*gin.RouterGroup
}

func NewEngine() *Engine {
	return &Engine{Engine: gin.Default()}
}

func handleFunc(handler func(c *Context)) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		handler(&Context{Context: c})
	}
}

func (server *Engine) Group(relativePath string, handlers ...func(c *Context)) *RouterGroup {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return &RouterGroup{server.Engine.Group(relativePath, RHandles...)}
}

func (server *Engine) GET(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return server.Engine.GET(relativePath, RHandles...)
}

func (server *Engine) POST(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return server.Engine.POST(relativePath, RHandles...)
}

func (r *RouterGroup) GET(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.GET(relativePath, rHandles...)
}

func (r *RouterGroup) POST(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.POST(relativePath, rHandles...)
}

func (r *RouterGroup) DELETE(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.DELETE(relativePath, rHandles...)
}

func (r *RouterGroup) PATCH(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.PATCH(relativePath, rHandles...)
}

func (r *RouterGroup) Use(middlewares ...func(c *Context)) gin.IRoutes {
	rMiddlewares := make([]gin.HandlerFunc, 0)
	for _, middleware := range middlewares {
		rMiddlewares = append(rMiddlewares, handleFunc(middleware))
	}
	return r.RouterGroup.Use(rMiddlewares...)
}

func (c *Context) Error(httpStatus int, msg string) {
	c.JSON(httpStatus, gin.H{"error": msg})
}

func (c *Context) ErrorInvalidToken() {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
}

func (c *Context) Ok(obj any) {
	c.JSON(http.StatusOK, obj)
}

func (c *Context) ParamInt(key string) int {
	keyValueStr := c.Param(key)
	keyValueInt, _ := strconv.Atoi(keyValueStr)
	return keyValueInt
}

func (c *Context) ErrorBad(msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": msg})
}

func (c *Context) GetUserId() (int, bool) {
	authorization := c.Request.Header.Get("Authorization")
	if len(authorization) == 0 {
		return 0, false
	}
	index := strings.IndexByte(authorization, ' ')
	if index < 0 {
		return 0, false
	}
	_type := authorization[:index]
	if strings.Compare(_type, "Bearer") != 0 {
		return 0, false
	}
	index = strings.LastIndexByte(authorization, ' ') + 1
	tokenStr := authorization[index:]

	token, _ := jwt.Parse(tokenStr, checkToken)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims["userId"].(float64); ok {
			return int(userId), true
		}
	}
	return 0, false
}

func checkToken(_ *jwt.Token) (interface{}, error) {
	return []byte("abc"), nil
}
