package main

import (
	"time"

	"github.com/AtaullinShamil/test_task_trood_ai_helpdesk/config"
	"github.com/AtaullinShamil/test_task_trood_ai_helpdesk/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Load"))
	}

	rmqClient, err := rabbitmq.NewClient(cfg.RabbitMQ.URL)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "NewClient"))
	}
	defer rmqClient.Close()

	if err := rmqClient.DeclareQueue("Helpdesk"); err != nil {
		logger.Fatal(errors.Wrap(err, "DeclareQueue Helpdesk"))
	}
	if err := rmqClient.DeclareQueue("HelpdeskResponse"); err != nil {
		logger.Fatal(errors.Wrap(err, "DeclareQueue HelpdeskResponse"))
	}
	responseMsgs, err := rmqClient.ConsumeMessages("HelpdeskResponse")
	if err != nil {
		logger.Fatal(errors.Wrap(err, "ConsumeMessages HelpdeskResponse"))
	}

	app := fiber.New()
	app.Get("/helpdesk", sendHandler(rmqClient, responseMsgs, logger))

	if err := app.Listen(":3000"); err != nil {
		logger.Fatal(errors.Wrap(err, "Listen"))
	}
}

func sendHandler(rmqClient *rabbitmq.RabbitMQClient, responseMsgs <-chan amqp.Delivery, logger *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		queries := c.Queries()
		msg, ok := queries["msg"]
		if !ok {
			return c.Status(fiber.StatusBadRequest).SendString("missing message")
		}

		correlationID := uuid.New().String()

		err := rmqClient.PublishMessage("Helpdesk", msg, correlationID)
		if err != nil {
			logger.Error(errors.Wrap(err, "PublishMessage"))
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to publish message")
		}

		response, err := waitForResponse(responseMsgs, correlationID)
		if err != nil {
			logger.Error(errors.Wrap(err, "waitForResponse"))
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to receive response")
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": string(response),
		})
	}
}

func waitForResponse(responseMsgs <-chan amqp.Delivery, correlationID string) ([]byte, error) {
	timeout := time.After(time.Minute)

	for {
		select {
		case message := <-responseMsgs:
			if message.CorrelationId == correlationID {
				return message.Body, nil
			}
		case <-timeout:
			return nil, errors.New("timeout waiting for response")
		}
	}
}
