package controllers

import (
	"net/http"
	"api/notification/src/notification/application"
	"github.com/gin-gonic/gin"
)


type NotificationController struct {
	service *application.NotificationService
}

func NewNotificationController(service *application.NotificationService) *NotificationController {
	return &NotificationController{service: service}
}

func (c *NotificationController) ReceiveNotification(ctx *gin.Context) {
	var notificationMessage struct {
		OrderMessage string `json:"order_message"`
	}

	if err := ctx.ShouldBindJSON(&notificationMessage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}


	c.service.ProcessOrder(notificationMessage.OrderMessage)


	ctx.JSON(http.StatusOK, gin.H{
		"message": "Notificación procesada con éxito",
	})
}
