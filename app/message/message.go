package message

import (
    "time"
)

// SocketMessage is sent by a connected device
type SocketMessage struct {
    Action      string      `json:"action"`
    Event       string      `json:"event"`
    Data        string      `json:"data"`
    CreatedAt   time.Time   `json:"created_at"`
}
