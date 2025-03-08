package routes

import (
	"api/notification/src/notification/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(router *gin.RouterGroup){
	getAllController := http.SetUpGetAllNotificationController()

	router.GET("/notifications", getAllController.GetAllNotification)
}