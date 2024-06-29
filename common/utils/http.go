package utils

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"qq-bot/common/global"
)

func Post(ife string, data interface{}) {
	// 将数据编码为JSON格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		global.GLog.Error("post", zap.Error(err))
		return
	}

	url := global.GConfig.QQBot.Address + ife
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		global.GLog.Error("post", zap.Error(err))
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	// 发起请求
	client := &http.Client{}

	_, err = client.Do(req)
	if err != nil {
		global.GLog.Error("post", zap.Error(err))
		return
	}
}
