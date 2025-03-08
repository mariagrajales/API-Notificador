package adapters

import (
	"api/notification/src/database"
	"api/notification/src/notification/domain/entities"
	"database/sql"
	"log"
)

type NotificationRepo struct {
	db *sql.DB
}

func NewNotificationRepo() (*NotificationRepo, error) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	return &NotificationRepo{db: db}, nil
}

func (r *NotificationRepo) Create(notification entities.Notification) error {
	_, err := r.db.Exec("INSERT INTO notifications (message) VALUES (?)", notification.Message)
	if err != nil {
		log.Println("Error al guardar la notificación:", err)
		return err
	}
	return nil
}

func (r *NotificationRepo) GetAll() ([]entities.Notification, error) {
	rows, err := r.db.Query("SELECT * FROM notifications")
	if err != nil {
		log.Println("Error al obtener las notificaciones:", err)
		return nil, err
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var n entities.Notification
		err := rows.Scan(&n.ID, &n.Message)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notifications = append(notifications, n)
	}
	return notifications, nil
}