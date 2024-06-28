package api

import (
	"github.com/gin-gonic/gin"
	"qq-bot/common/global"
	"qq-bot/common/utils"
	"qq-bot/model/common/req"
	"qq-bot/model/common/resp"
)

type EventApi struct{}

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

	//TODO 复读逻辑 控制概率大概在1/15左右

	//TODO 阴阳怪气逻辑 控制概率大概在1/20左右

	// 如果第一个消息数组第一个是@类型 且 @对象是机器人 才进行回复
	if msg.Message[0].Type == "at" && msg.Message[0].Data.Qq == global.GConfig.QQBot.Qq {
		// 如果只进行了@, 而没有内容, 则回复一个简短的问号
		// 如果消息数组长度为2, 则进行AI回复
		if len(msg.Message) == 1 {
			resp.ReplyOk(c, "？")
		} else if len(msg.Message) == 2 {
			if msg.Message[1].Type == "text" {

				resp.ReplyOk(c, msg.Message[1].Data.Text)
			}
		}
	}
}
