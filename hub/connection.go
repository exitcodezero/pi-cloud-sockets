package hub

import (
	"github.com/exitcodezero/picloud/message"
	"github.com/pborman/uuid"
)

// Connection maintains info about the connected device
// and subscribed events
type Connection struct {
	ID         string
	IPAddress  string
	Subscribed []string
	Out        chan message.SocketMessage
}

// NewConnection constructs a new Connection
func NewConnection(ipAddress string) Connection {
	hc := Connection{}
	hc.ID = uuid.New()
	hc.IPAddress = ipAddress
	hc.Out = make(chan message.SocketMessage)
	return hc
}
