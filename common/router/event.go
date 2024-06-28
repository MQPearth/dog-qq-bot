package router

import (
	"github.com/gin-gonic/gin"
	"qq-bot/api"
)

type EventRouter struct{}

func (s *EventRouter) InitPublicRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("event")
	eventApi := api.ApiGroup.EventApi
	authorityRouter.POST("post", eventApi.Post)
}

func (s *EventRouter) InitPrivateRouter(Router *gin.RouterGroup) {

}
