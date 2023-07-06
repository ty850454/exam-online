package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DebugMode bool
	DB        *gorm.DB
	Router    *gin.Engine
)
