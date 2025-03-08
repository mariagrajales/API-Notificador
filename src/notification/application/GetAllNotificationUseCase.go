package application

import "api/notification/src/notification/domain/ports"

type GetAllNotificationUseCase struct {
	NotificationService ports.NotificationRepository
}

func NewGetAllNotificationUseCase(service ports.NotificationRepository) *GetAllNotificationUseCase {
	return &GetAllNotificationUseCase{NotificationService: service}
}

func (s *GetAllNotificationUseCase) GetAll() ([]string, error) {
	notifications, err := s.NotificationService.GetAll()
	if err != nil {
		return nil, err
	}

	var messages []string
	for _, notification := range notifications {
		messages = append(messages, notification.Message)
	}

	return messages, nil
}