package main

import (
	"github.com/AtaullinShamil/test_task_trood_ai_helpdesk/config"
	"github.com/AtaullinShamil/test_task_trood_ai_helpdesk/pkg/rabbitmq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

	messages, err := rmqClient.ConsumeMessages("Helpdesk")
	if err != nil {
		logger.Fatal(errors.Wrap(err, "ConsumeMessages"))
	}

	logger.Info("Service started")

	for message := range messages {
		// To Do
		responseMessage := "Processed: " + string(message.Body)

		err := rmqClient.PublishMessage("HelpdeskResponse", responseMessage, message.CorrelationId)
		if err != nil {
			logger.Error(errors.Wrap(err, "PublishMessage"))
		}
	}
}
