package hub

import (
	"github.com/exitcodezero/picloud/message"
	"github.com/pborman/uuid"
	"time"
)

// Connection maintains info about the connected device
// and subscribed events
type Connection struct {
	ID          string
	ClientName 	string
	IPAddress   string
	ConnectedAt time.Time
	Subscribed  []string
	Out         chan message.SocketMessage
}

// NewConnection constructs a new Connection
func NewConnection(ipAddress string, clientName string) Connection {
	hc := Connection{}
	hc.ID = uuid.New()
	hc.ClientName = clientName
	hc.IPAddress = ipAddress
	hc.ConnectedAt = time.Now().UTC()
	hc.Out = make(chan message.SocketMessage)
	return hc
}
