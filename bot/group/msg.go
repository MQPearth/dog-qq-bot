package group

import (
	"qq-bot/common/utils"
	"qq-bot/model/common/req"
	"strconv"
)

// SendGroupMsg 发送群聊消息
func SendGroupMsg(id int, content string) {

	msg := req.Message{
		Data: req.Data{Text: content},
		Type: "text",
	}

	data := map[string]interface{}{
		"group_id": id,
		"message":  msg,
	}
	utils.Post("/send_group_msg_async", data)
}

// SendAtGroupMsg 发送群聊艾特消息
func SendAtGroupMsg(id int, userId int, content string) {
	var msgArr [2]req.Message
	msgArr[0] = req.Message{
		Data: req.Data{Qq: strconv.Itoa(userId)},
		Type: "at",
	}

	msgArr[1] = req.Message{
		Data: req.Data{Text: " " + content},
		Type: "text",
	}

	data := map[string]interface{}{
		"group_id": id,
		"message":  msgArr,
	}
	utils.Post("/send_group_msg_async", data)
}
