package queue

import (
	"fmt"
	"github.com/streadway/amqp"
)

type MessageQueue struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func InitMessageQueue() *MessageQueue {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", err)
		return nil
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel:", err)
		return nil
	}

	return &MessageQueue{conn: conn, ch: ch}
}

func (mq *MessageQueue) PushImageURLs(imageURLs []string) error {
	q, err := mq.ch.QueueDeclare(
		"image_processing_queue", // name of the queue
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		return err
	}

	for _, url := range imageURLs {
		err = mq.ch.Publish(
			"",           // exchange
			q.Name,       // routing key (queue name)
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(url),
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}
