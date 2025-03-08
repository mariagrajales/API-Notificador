package controllers

import (
	"api/notification/src/notification/application"

	"github.com/gin-gonic/gin"
)

type GetAllNotificationController struct {
	NotificationService *application.GetAllNotificationUseCase
}

func NewGetAllNotificationController(service *application.GetAllNotificationUseCase) *GetAllNotificationController {
	return &GetAllNotificationController{NotificationService: service}
}


func (ctr *GetAllNotificationController) GetAllNotification(ctx *gin.Context) {
	messages, err := ctr.NotificationService.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Error al obtener las notificaciones",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"messages": messages,
	})
}
