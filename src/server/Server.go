package server

import (
	"log"

	"api/notification/src/database"
	"api/notification/src/notification/application"
	"api/notification/src/notification/infraestructure/adapters"
	"api/notification/src/notification/infraestructure/http/routes"
    "api/notification/src/config" 
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	http     string
	port     string
	httpAddr string
}

func NewServer(http, port string) Server {
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		engine:   gin.Default(),
		http:     http,
		port:     port,
		httpAddr: http + ":" + port,
	}

	database.Connect()

	srv.engine.RedirectTrailingSlash = true
	srv.registerRoutes()

	notificationRepository, err := adapters.NewNotificationRepo()
	if err != nil {
		log.Fatalf("Error al crear el repositorio de notificaciones: %v", err)
	}

	notificationService := application.NewNotificationService(notificationRepository)

	// Aquí es donde se debe llamar explícitamente a ConsumeOrders
	consumer := adapters.NewRabbitMQConsumer(notificationService)
	go consumer.ConsumeOrders() // Ejecutamos ConsumeOrders de forma concurrente

	srv.engine.Use(config.ConfigurationCors())

	return srv
}

func (s *Server) registerRoutes() {
	notificationRoutes := s.engine.Group("/api")

	routes.NotificationRoutes(notificationRoutes)

}

func (s *Server) Run() error {
	log.Println("Server running on " + s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

