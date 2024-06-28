
## 狗群友bot

功能

1. AI问答 (@机器人)
2. 半夜发鬼故事并@一个群友
3. 在群友聊天时, 6.6%的概率进行复读
4. 在群友聊天时, 5%的概率进行阴阳
5. 每天早晨发毒鸡汤

基于[NapCatQQ](https://github.com/NapNeko/NapCatQQ)

NapCatQQ配置

```json
{
  "http": {
    "enable": true,
    "host": "0.0.0.0",
    "port": 3001,
    "secret": "",
    "enableHeart": false,
    "enablePost": true,
    "postUrls": ["http://192.168.93.1:19997/api/event/post"]
  },
  "ws": {
    "enable": false,
    "host": "",
    "port": 3001
  },
  "reverseWs": {
    "enable": false,
    "urls": []
  },
  "debug": false,
  "heartInterval": 60000,
  "messagePostFormat": "array",
  "enableLocalFile2Url": false,
  "musicSignUrl": "",
  "reportSelfMessage": false,
  "token": "",
  "GroupLocalTime": {
    "Record": false,
    "RecordList": []
  }
}

```