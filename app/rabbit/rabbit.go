package rabbit

import (
    "log"
    "github.com/streadway/amqp"
    "app/config"
)

// Conn
var Conn *amqp.Connection

func init() {
    rabbitURL := config.RabbitURL()
    var err error
    Conn, err = amqp.Dial(rabbitURL)
    if err != nil {
		log.Fatal(err)
	}
}
