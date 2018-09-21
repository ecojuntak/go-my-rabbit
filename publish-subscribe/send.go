package main

import (
	"log"

	"github.com/streadway/amqp"
)

const EXCHANGE_NAME = "logs"

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

func declareExchange(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		EXCHANGE_NAME, // name
		"fanout",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")
}

func publishMessage(ch *amqp.Channel, body string) (err error) {
	err = ch.Publish(
		EXCHANGE_NAME, // exchange
		"",            // routing key
		false,         // mandatory
		false,         // immediate
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

	declareExchange(ch)

	body := "Hello World!"
	err := publishMessage(ch, body)

	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
