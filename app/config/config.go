package config

import (
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

// Port is the env variable PORT
var Port string

func init()  {
    if os.Getenv("PORT") != "" {
        Port = os.Getenv("PORT")
    } else {
        Port = "9000"
    }
}
