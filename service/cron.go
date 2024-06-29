package service

import (
	"go.uber.org/zap"
	"qq-bot/bot/group"
	"qq-bot/common/cons"
	"qq-bot/common/global"
	"qq-bot/common/utils"
	"strconv"
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
	content := ServiceGroup.DbService.RandomGetOne(cons.GhostStory)
	for _, id := range global.GConfig.QQBot.ActiveGroup {
		botId, _ := strconv.ParseInt(global.GConfig.QQBot.Qq, 10, 64)
		list := group.GetGroupMemberListRemoveSelf(id, botId)
		length := len(list)
		if length == 0 {
			continue
		}
		userId := list[utils.RandomInt(length)]
		group.SendAtGroupMsg(id, userId, content)
	}
}

// SendInspirationalStory 早上毒鸡汤
func (c *CronService) SendInspirationalStory() {
	defer func() {
		if err := recover(); err != nil {
			global.GLog.Error("SendInspirationalStory", zap.Any("info", err))
		}
	}()
	content := ServiceGroup.DbService.RandomGetOne(cons.InspirationalStory)
	for _, id := range global.GConfig.QQBot.ActiveGroup {
		group.SendGroupMsg(id, content)
	}
}
