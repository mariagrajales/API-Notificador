package http

import (
	"api/notification/src/notification/application"
	"api/notification/src/notification/domain/ports"
	"api/notification/src/notification/infraestructure/adapters"
	"api/notification/src/notification/infraestructure/http/controllers"
	"log"
)


var (
	notificationRepository ports.NotificationRepository
)

func init() {
	var err error
	notificationRepository, err = adapters.NewNotificationRepo()

	if err != nil {
		log.Fatalf("Error al crear el repositorio de notificaciones: %v", err)
	}
}


func SetUpNotificationController() *controllers.NotificationController {
	service := application.NewNotificationService(notificationRepository)
	return controllers.NewNotificationController(service)
}


func SetUpRabbitMQConsumer() *adapters.RabbitMQConsumer {
	service := application.NewNotificationService(notificationRepository)
	return adapters.NewRabbitMQConsumer(service)
}

func SetUpGetAllNotificationController() *controllers.GetAllNotificationController {
	service := application.NewGetAllNotificationUseCase(notificationRepository)
	return controllers.NewGetAllNotificationController(service)
}