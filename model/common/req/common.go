package req

type BaseData struct {
	SelfID        int    `json:"self_id"`
	Time          int    `json:"time"`
	PostType      string `json:"post_type"`
	MessageFormat string `json:"message_format"`
	MessageType   string `json:"message_type"`
}

type MsgData struct {
	BaseData
	UserID     int       `json:"user_id"`
	MessageID  int       `json:"message_id"`
	MessageSeq int       `json:"message_seq"`
	RealID     int       `json:"real_id"`
	Sender     Sender    `json:"sender"`
	RawMessage string    `json:"raw_message"`
	Font       int       `json:"font"`
	SubType    string    `json:"sub_type"`
	Message    []Message `json:"message"`
	GroupID    int       `json:"group_id"`
}
type Sender struct {
	UserID   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Role     string `json:"role"`
}
type Data struct {
	Text     string `json:"text"`
	Qq       string `json:"qq"`
	File     string `json:"file"`
	URL      string `json:"url"`
	FileSize string `json:"file_size"`
}

type Message struct {
	Data    Data   `json:"data,omitempty"`
	Type    string `json:"type"`
	SubType int    `json:"subType,omitempty"`
}
