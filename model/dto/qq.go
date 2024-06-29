package dto

type BaseResp struct {
	RetCode int          `json:"retCode"`
	Data    []MemberInfo `json:"data"`
}

type MemberInfo struct {
	UserId int64 `json:"user_id"`
}
