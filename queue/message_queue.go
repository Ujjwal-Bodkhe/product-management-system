package queue

import (
	"github.com/streadway/amqp"
	"log"
)

type MessageQueue struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewMessageQueue(url string) *MessageQueue {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}

	return &MessageQueue{conn, ch}
}

func (mq *MessageQueue) PushImageURLs(imageURLs []string) {
	for _, url := range imageURLs {
		err := mq.ch.Publish(
			"",        // exchange
			"image_queue", // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(url),
			},
		)
		if err != nil {
			log.Println("Failed to send message:", err)
		}
	}
}
