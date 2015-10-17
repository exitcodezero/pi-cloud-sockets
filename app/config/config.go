package config

import (
    "fmt"
    "os"
)

// APIKey is the env variable API_KEY
var APIKey = os.Getenv("API_KEY")

// UseTSL is the env variable USE_TSL
var UseTSL = os.Getenv("USE_TSL")

// KeyFile is the env variable KEY_FILE
var KeyFile = os.Getenv("KEY_FILE")

// CertFile is the env variable CERT_FILE
var CertFile = os.Getenv("CERT_FILE")

// RabbitURL returns the env variable RABBIT_URL if present.
// If not it returns a local RabbitMQ URL for docker-compose
func RabbitURL() string  {
    rabbitURL := os.Getenv("RABBIT_URL")
    if rabbitURL == "" {
        composeIP := os.Getenv("RABBIT_PORT_5672_TCP_ADDR")
        rabbitURL = fmt.Sprintf("amqp://guest:guest@%s:5672", composeIP)
    }
    return rabbitURL
}
