package group

import (
	"qq-bot/common/utils"
	"qq-bot/model/dto"
)

func GetGroupMemberListRemoveSelf(id int64, selfId int64) []int64 {
	data := map[string]interface{}{
		"group_id": id,
	}
	resp := utils.Post("/get_group_member_list", data)

	var result dto.BaseResp
	utils.ByteToStruct([]byte(resp), &result)

	var userIds []int64

	if result.RetCode != 0 {
		return userIds
	}

	for _, member := range result.Data {
		if member.UserId == selfId {
			continue
		}
		userIds = append(userIds, member.UserId)
	}
	return userIds
}
