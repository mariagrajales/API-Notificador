package application

import (
	"api/notification/src/notification/domain/entities"
	"api/notification/src/notification/domain/ports"
	"log"
)

type NotificationService struct {
	repo ports.NotificationRepository
}

func NewNotificationService(repo ports.NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) ProcessOrder(orderMessage string) {
	notification := entities.Notification{Message: "Orden recibida: " + orderMessage}
	err := s.repo.Create(notification)
	if err != nil {
		log.Println("Error al procesar la orden:", err)
	}
}
