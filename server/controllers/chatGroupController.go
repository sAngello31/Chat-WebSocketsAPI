package controllers

import (
	"chat_websocket/models"

	"github.com/gin-gonic/gin"
)

type ChatGroupController struct {
	ChatGroupRepo *models.ChatGroupRepository
}

var ThisChatGroupController *ChatGroupController

func NewChatGroupController(chatGroup *models.ChatGroupRepository) {
	ThisChatGroupController = &ChatGroupController{ChatGroupRepo: chatGroup}
}

func (ctrl *ChatGroupController) GetAllChatsGroups(c *gin.Context) {
	ctrl.ChatGroupRepo.GetAllChatsGroups(c)
}

func (ctrl *ChatGroupController) InsertChatGroup(c *gin.Context) {
	ctrl.ChatGroupRepo.InsertChatGroup(c)
}
