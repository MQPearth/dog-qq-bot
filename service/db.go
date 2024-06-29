package service

import (
	"go.uber.org/zap"
	"qq-bot/common/global"
	"qq-bot/common/utils"
	"qq-bot/model/po"
)

type DbService struct{}

func (b *DbService) RandomGetOne(typ int) string {

	length := len(global.GStory[typ])

	id := global.GStory[typ][utils.RandomInt(length)]

	var story string

	err := global.GDb.Raw("SELECT `content`  FROM story WHERE id = ?", id).Scan(&story).Error

	if err != nil {
		global.GLog.Error("RandomGetOne", zap.Error(err))
		return ""
	}

	return story
}

func (b *DbService) QueryAllStoryIdAndType() map[int][]int {

	typeMap := make(map[int][]int)

	var stories []po.BaseStory
	err := global.GDb.Raw("SELECT `id`, `type` FROM story WHERE `status` = 1").Scan(&stories).Error
	if err != nil {
		global.GLog.Error("QueryAllStoryIdAndType", zap.Error(err))
		return typeMap
	}

	for _, story := range stories {
		typeMap[story.Type] = append(typeMap[story.Type], story.Id)
	}

	return typeMap
}
