package queue

import (
	"fmt"
	"github.com/streadway/amqp"
)

type MessageQueue struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func InitQueue() (*MessageQueue, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MessageQueue{connection: conn, channel: channel}, nil
}

func (q *MessageQueue) PushImageURLs(urls []string) error {
	for _, url := range urls {
		err := q.channel.Publish(
			"",     // exchange
			"image_queue", // routing key
			false,  // mandatory
			false,  // immediate
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
