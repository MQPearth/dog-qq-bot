package global

import (
	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"qq-bot/common/config"
)

var (
	GConfig config.Config
	GLog    *zap.Logger
	GDb     *gorm.DB
	GAi     *qianfan.ChatCompletion
	GCron   *cron.Cron
	GStory  map[int][]int
)
