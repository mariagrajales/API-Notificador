package adapters

import (
	"api/notification/src/notification/application"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	service *application.NotificationService
}

func NewRabbitMQConsumer(service *application.NotificationService) *RabbitMQConsumer {
	return &RabbitMQConsumer{service: service}
}

func (c *RabbitMQConsumer) ConsumeOrders() {
	conn, err := amqp.Dial("amqp://user:password@54.235.169.219:5672")
	if err != nil {
		log.Fatal("Error al conectar con RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error al abrir un canal:", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"order_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error al declarar exchange:", err)
	}

	q, err := ch.QueueDeclare("", false, true, true, false, nil)
	if err != nil {
		log.Fatal("Error al declarar cola:", err)
	}
	log.Printf("Cola temporal declarada con nombre: %s\n", q.Name)

	err = ch.QueueBind(
		q.Name,
		"orden_topic",
		"order_topic",
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error al enlazar cola con exchange:", err)
	}
	log.Printf("Cola %s vinculada con exchange 'orders_exchange' usando la clave de enrutamiento 'orden_topic'.\n", q.Name)

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Error al consumir mensajes:", err)
	}

	log.Println("Esperando órdenes...")

	for msg := range msgs {
		log.Printf("Mensaje recibido: %s\n", string(msg.Body))
		c.service.ProcessOrder(string(msg.Body))
	}
}
