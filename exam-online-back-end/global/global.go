package global

import (
	"gorm.io/gorm"
)

var (
	DebugMode  bool
	DB         *gorm.DB
	ApiRouter  *RouterGroup
	UserRouter *RouterGroup
)
