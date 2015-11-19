package socket

import (
	"github.com/exitcodezero/picloud/hub"
	"github.com/exitcodezero/picloud/message"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func writeSocket(socket *websocket.Conn, c hub.Connection) {
	for {
		m := <-c.Out
		socket.WriteJSON(&m)
	}
}

// Handler handles websocket connections at /connect
func Handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer socket.Close()

	c := hub.NewConnection(socket.RemoteAddr().String())
	defer hub.UnsubscribeAll(c)

	go writeSocket(socket, c)

	for {
		m := message.SocketMessage{}
		m.CreatedAt = time.Now().UTC()

		err = socket.ReadJSON(&m)
		if err != nil {
			break
		}

		switch m.Action {
		case "publish":
			hub.Publish(m)
		case "subscribe":
			hub.Subscribe(m.Event, c)
		case "unsubscribe":
			hub.Unsubscribe(m.Event, c)
		case "unsubscribe:all":
			hub.UnsubscribeAll(c)
		}
	}
}
