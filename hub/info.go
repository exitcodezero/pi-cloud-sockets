package hub

import (
	"time"
)

type connectionInfo struct {
	ClientName  string    `json:"client_name"`
	IPAddress   string    `json:"ip_address"`
	ConnectedAt time.Time `json:"connected_at"`
}

type eventInfo struct {
	Name        string           `json:"name"`
	Connections []connectionInfo `json:"connections"`
}

type infoMessage struct {
	Subscriptions  []eventInfo      `json:"subscriptions"`
	AllConnections []connectionInfo `json:"all_connections"`
	CreatedAt      time.Time        `json:"created_at"`
}
