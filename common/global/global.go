package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"qq-bot/common/config"
)

var (
	GConfig config.Config
	GLog    *zap.Logger
	GDb     *gorm.DB
)
