package config

import (
    "os"
)

// APIKey is the env variable API_KEY
var APIKey = os.Getenv("API_KEY")
