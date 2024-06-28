package api

import "qq-bot/service"

type ApiGroupEnter struct {
	EventApi
}

var (
	// ApiGroup public
	ApiGroup = new(ApiGroupEnter)

	// dbService private
	dbService = service.ServiceGroup.DbService
)
