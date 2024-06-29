package api

import (
	"context"
	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"qq-bot/bot/group"
	"qq-bot/common/global"
	"qq-bot/common/utils"
	"qq-bot/model/common/req"
	"qq-bot/model/common/resp"
)

type EventApi struct{}

func aiMsg(content string) string {
	response, err := global.GAi.Do(
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage(content),
			},
		},
	)
	if err != nil {
		global.GLog.Error("模型调用失败", zap.Error(err))
		return "我寄了"
	}
	return response.Result
}

// Post 事件
func (a *EventApi) Post(c *gin.Context) {
	rawData, err := c.GetRawData()
	if err != nil {
		resp.EmptyOk(c)
		return
	}
	var data req.BaseData
	utils.ByteToJson(rawData, &data)

	// 通知类型不是消息的不处理
	if data.PostType != "message" {
		resp.EmptyOk(c)
		return
	}
	// 消息格式不是数组的不处理
	if data.MessageFormat != "array" {
		resp.EmptyOk(c)
		return
	}
	// 不是群聊消息不处理
	if data.MessageType != "group" {
		resp.EmptyOk(c)
		return
	}

	var msg req.MsgData
	utils.ByteToJson(rawData, &msg)

	if !global.GConfig.QQBot.IsActive(msg.GroupID) {
		resp.EmptyOk(c)
		return
	}

	// 如果第一个消息数组第一个是@类型 且 @对象是机器人 才进行回复
	if msg.Message[0].Type == "at" && msg.Message[0].Data.Qq == global.GConfig.QQBot.Qq {
		// 如果只进行了@, 而没有内容, 则回复一个简短的问号
		// 如果消息数组长度为2, 且第二个消息段是文本 则进行AI回复
		if len(msg.Message) == 1 {
			resp.ReplyOk(c, "？")
			return
		} else if len(msg.Message) == 2 {
			if msg.Message[1].Type == "text" {
				resp.ReplyOk(c, aiMsg(msg.Message[1].Data.Text))
				return
			}
		}
	}

	if msg.Message[0].Type == "text" {
		random := utils.RandomInt(100)
		global.GLog.Debug("随机数字: ", zap.Int("random", random))

		if random >= 0 && random < 6 {
			// 复读
			group.SendGroupMsg(msg.GroupID, msg.Message[0].Data.Text)
		} else if random >= 52 && random < 57 {
			// 阴阳
			content := dbService.RandomGetOne(3)
			group.SendGroupMsg(msg.GroupID, content)
		} else if random >= 94 && random < 100 {
			// ?
			group.SendGroupMsg(msg.GroupID, "?")
		}
	}

	resp.EmptyOk(c)
	return
}
