package service

import (
	"go.uber.org/zap"
	"qq-bot/bot/group"
	"qq-bot/common/global"
)

type CronService struct{}

// LoadStory 一分钟加载一次数据库
func (c *CronService) LoadStory() {
	defer func() {
		if err := recover(); err != nil {
			global.GLog.Error("LoadStory", zap.Any("info", err))
		}
	}()
	global.GStory = ServiceGroup.DbService.QueryAllStoryIdAndType()
}

// SendGhostStories 半夜发送鬼故事
func (c *CronService) SendGhostStories() {
	defer func() {
		if err := recover(); err != nil {
			global.GLog.Error("SendGhostStories", zap.Any("info", err))
		}
	}()

}

// SendInspirationalStory 早上毒鸡汤
func (c *CronService) SendInspirationalStory() {
	defer func() {
		if err := recover(); err != nil {
			global.GLog.Error("SendInspirationalStory", zap.Any("info", err))
		}
	}()
	content := ServiceGroup.DbService.RandomGetOne(2)
	for _, id := range global.GConfig.QQBot.ActiveGroup {
		group.SendGroupMsg(id, content)
	}
}
