package rabbitmq

import (
	"log"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewClient(url string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, errors.Wrap(err, "Dial")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Channel")
	}

	return &RabbitMQClient{
		Connection: conn,
		Channel:    ch,
	}, nil
}

func (client *RabbitMQClient) DeclareQueue(queueName string) error {
	_, err := client.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.Wrapf(err, "QueueDeclare for %s", queueName)
	}
	return nil
}

func (client *RabbitMQClient) PublishMessage(queueName, message, correlationID string) error {
	if client.Channel == nil || client.Connection.IsClosed() {
		return errors.New("RabbitMQ connection or channel is closed")
	}

	err := client.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          []byte(message),
			CorrelationId: correlationID,
		},
	)
	if err != nil {
		return errors.Wrapf(err, "Publish to %s", queueName)
	}
	return nil
}

func (client *RabbitMQClient) ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := client.Channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "Consume from %s", queueName)
	}
	return msgs, nil
}

func (client *RabbitMQClient) Close() {
	if err := client.Channel.Close(); err != nil {
		log.Println("Failed to close channel:", err)
	}
	if err := client.Connection.Close(); err != nil {
		log.Println("Failed to close connection:", err)
	}
}
