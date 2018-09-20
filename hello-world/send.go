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

func publishMessage(ch *amqp.Channel, q amqp.Queue, body string) (err error) {
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	return
}

func main() {
	conn := connectToServer()
	defer conn.Close()

	ch := openChannel(conn)
	defer ch.Close()

	q := declareQueue(ch)

	body := "Hello World!"
	err := publishMessage(ch, q, body)

	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
