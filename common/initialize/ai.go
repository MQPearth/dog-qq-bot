package initialize

import (
	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
	"qq-bot/common/global"
)

func Ai() {
	qianfan.GetConfig().AccessKey = global.GConfig.Ai.AccessKey
	qianfan.GetConfig().SecretKey = global.GConfig.Ai.SecretKey

	// Yi-34B-Chat暂时免费
	global.GAi = qianfan.NewChatCompletion(qianfan.WithModel("Yi-34B-Chat"))

	// 支持的模型
	list := global.GAi.ModelList()
	for _, s := range list {
		global.GLog.Debug(s)
	}
}