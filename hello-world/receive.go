package main

import (
	"log"

	"github.com/streadway/amqp"
)

const QUEUE_NAME = "hello"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectToServer() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func openChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

func declareQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		QUEUE_NAME, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return q
}

func main() {
	conn := connectToServer()
	defer conn.Close()

	ch := openChannel(conn)
	defer ch.Close()

	q := declareQueue(ch)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
