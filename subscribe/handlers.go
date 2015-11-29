package subscribe

import (
	"github.com/exitcodezero/picloud/hub"
	"github.com/exitcodezero/picloud/message"
	"github.com/gorilla/context"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func writeSocket(socket *websocket.Conn, c *hub.Connection) {
	for {
		m := <-c.Out
		socket.WriteJSON(&m)
	}
}

// Handler handles websocket connections at /connect
func Handler(w http.ResponseWriter, r *http.Request) {

	clientName, _ := context.Get(r, "ClientName").(string)
	context.Clear(r)

	// Upgrade the request
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer socket.Close()

	// Create a Connection instance
	c := hub.NewConnection(socket.RemoteAddr().String(), clientName)
	hub.Manager.RegisterConnection(&c)
	defer hub.Manager.Cleanup(&c)

	// Start pushing outbound messages from a goroutine
	go writeSocket(socket, &c)

	// Handle inbound subscription messages
	for {
		m := message.SocketMessage{}
		m.CreatedAt = time.Now().UTC()

		err = socket.ReadJSON(&m)
		if err != nil {
			break
		}

		switch m.Action {
		case "subscribe":
			hub.Manager.Subscribe(m.Event, &c)
		case "unsubscribe":
			hub.Manager.Unsubscribe(m.Event, &c)
		case "unsubscribe:all":
			hub.Manager.UnsubscribeAll(&c)
		}
	}
}
