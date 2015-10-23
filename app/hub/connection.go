package hub

import (
    "github.com/pborman/uuid"
    "app/message"
)

// Connection maintains info about the connected device
// and subscribed events
type Connection struct {
    ID          string
    Subscribed  []string
    Out         chan message.SocketMessage
}

// NewConnection constructs a new Connection
func NewConnection() Connection {
    hc := Connection{}
    hc.ID = uuid.New()
    hc.Out = make(chan message.SocketMessage)
    return hc
}
