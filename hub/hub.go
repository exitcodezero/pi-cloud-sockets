package hub

import (
	"github.com/exitcodezero/picloud/message"
)

type manager struct {
	ToPublish chan message.SocketMessage
	Subscribed map[string][]*Connection
	Connections []*Connection
}

func (m *manager) ProcessSubscriptions() {
	for {
		message := <- m.ToPublish
		for _, c := range m.Subscribed[message.Event] {
			c.Out <- message
		}
	}
}

func (m *manager) RegisterConnection(c *Connection) {
	m.Connections = append(m.Connections, c)
}

func (m *manager) UnregisterConnection(c *Connection) {
	i := findConnectionIndex(c, m.Connections)
	if i != -1 {
		m.Connections = append(m.Connections[:i], m.Connections[i+1:]...)
	}
	m.Connections = append(m.Connections, c)
}

func (m *manager) Cleanup(c *Connection) {
	m.UnregisterConnection(c)
	m.UnsubscribeAll(c)
}

func (m *manager) Publish(msg message.SocketMessage) {
	m.ToPublish <- msg
}

func (m *manager) Subscribe(event string, c *Connection) {
	i := findConnectionIndex(c, m.Subscribed[event])
	if i == -1 {
		m.Subscribed[event] = append(m.Subscribed[event], c)
	}
}

func (m *manager) Unsubscribe(event string, c *Connection) {
	i := findConnectionIndex(c, m.Subscribed[event])
	if i != -1 {
		m.Subscribed[event] = append(m.Subscribed[event][:i], m.Subscribed[event][i+1:]...)
	}
}

func (m *manager) UnsubscribeAll(c *Connection) {
	for e := range m.Subscribed {
		m.Unsubscribe(e, c)
	}
}

func (m *manager) Info() infoMessage {
	im := infoMessage{
		Subscriptions: m.eventInfoSlice(),
		AllConnections: m.connectionInfoSlice(),
	}
	return im
}

func (m *manager) eventInfoSlice() []eventInfo {
	e := make([]eventInfo, 0)
	for k, connections := range m.Subscribed {
		ev := eventInfo{}
		ev.Name = k
		for _, c := range connections {
			cInfo := connectionInfo{
				IPAddress:   c.IPAddress,
				ConnectedAt: c.ConnectedAt,
			}
			ev.Connections = append(ev.Connections, cInfo)
		}
		e = append(e, ev)
	}
	return e
}

func (m *manager) connectionInfoSlice() []connectionInfo {
	var ci []connectionInfo
	for _, c := range m.Connections {
		cInfo := connectionInfo{
			IPAddress:   c.IPAddress,
			ConnectedAt: c.ConnectedAt,
		}
		ci = append(ci, cInfo)
	}
	return ci
}

// Manager controls all publish/subscribe actions for connections
var Manager manager

func init() {

	Manager = manager{
		ToPublish: make(chan message.SocketMessage),
		Subscribed: make(map[string][]*Connection),
		Connections: make([]*Connection, 0),
	}

	go Manager.ProcessSubscriptions()
}
