package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"qq-bot/common/global"
)

func Post(ife string, data interface{}) string {
	// 将数据编码为JSON格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		global.GLog.Error("post", zap.Error(err))
		return ""
	}

	url := global.GConfig.QQBot.Address + ife
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		global.GLog.Error("post", zap.Error(err))
		return ""
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	// 发起请求
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		global.GLog.Error("post", zap.Error(err))
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			global.GLog.Error("Body.Close", zap.Error(err))
		}
	}(resp.Body)

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应错误:", err)
		return ""
	}
	return string(body)
}
