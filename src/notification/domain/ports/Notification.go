package ports

import "api/notification/src/notification/domain/entities"

type NotificationRepository interface {
	Create(notification entities.Notification) error
	GetAll() ([]entities.Notification, error)
}
