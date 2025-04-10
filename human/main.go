package main

import (
	"bufio"
	"fmt"
	"os"

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

	if err := rmqClient.DeclareQueue("HelpdeskHuman"); err != nil {
		logger.Fatal(errors.Wrap(err, "DeclareQueue HelpdeskHuman"))
	}
	if err := rmqClient.DeclareQueue("HelpdeskResponse"); err != nil {
		logger.Fatal(errors.Wrap(err, "DeclareQueue HelpdeskResponse"))
	}

	messages, err := rmqClient.ConsumeMessages("HelpdeskHuman")
	if err != nil {
		logger.Fatal(errors.Wrap(err, "ConsumeMessages HelpdeskHuman"))
	}

	logger.Info("Human response service started")

	for message := range messages {
		fmt.Println("New message from Helpdesk queue:")
		fmt.Println(string(message.Body))

		fmt.Println("Enter your response: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		response := scanner.Text()

		err = rmqClient.PublishMessage("HelpdeskResponse", response, message.CorrelationId)
		if err != nil {
			logger.Error(errors.Wrap(err, "PublishMessage to HelpdeskResponse"))
		} else {
			logger.Info("Response sent to HelpdeskResponse queue")
		}
	}
}
