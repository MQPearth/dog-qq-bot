package service

import (
	"qq-bot/common/global"
)

type CronService struct{}

// LoadStory 一分钟加载一次数据库
func (c *CronService) LoadStory() {
	global.GStory = ServiceGroup.DbService.QueryAllStoryIdAndType()
}

// SendGhostStories 半夜发送鬼故事
func (c *CronService) SendGhostStories() {

}

// SendInspirationalStory 早上毒鸡汤
func (c *CronService) SendInspirationalStory() {

}
