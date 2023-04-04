package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func sendToJs() {

	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queueName := "go-to-js"
	exchange := "contactGo"

	// assert exchange
	err = ch.ExchangeDeclare(exchange, "fanout", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// bind
	err = ch.QueueBind(q.Name, "", exchange, false, nil)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}

	fmt.Printf("Waiting for messages on %s\n", queueName)

	msg := "Hello from Go!"
	err = ch.Publish(exchange, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	fmt.Println("Message sent")

}

func receivedFromJs() {

	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queueName := "js-to-go"
	exchange := "contactJS"

	// assert exchange
	err = ch.ExchangeDeclare(exchange, "fanout", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func main() {
	sendToJs()
	receivedFromJs()
}
