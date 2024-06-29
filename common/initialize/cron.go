package initialize

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"qq-bot/common/global"
	"qq-bot/service"
)

func Cron() {
	// 创建一个新的cron实例
	global.GCron = cron.New()

	_, err := global.GCron.AddFunc("*/1 * * * *", service.ServiceGroup.CronService.LoadStory)

	if err != nil {
		global.GLog.Error("LoadStory", zap.Error(err))
		return
	}

	_, err = global.GCron.AddFunc("30 7 * * *", service.ServiceGroup.CronService.SendInspirationalStory)

	if err != nil {
		global.GLog.Error("SendInspirationalStory", zap.Error(err))
		return
	}

	_, err = global.GCron.AddFunc("49 21 * * *", service.ServiceGroup.CronService.SendGhostStories)

	if err != nil {
		global.GLog.Error("SendGhostStories", zap.Error(err))
		return
	}

	global.GCron.Start()
	global.GLog.Info("定时任务初始化完成")
}
