package socket

import (
    "time"
    "net/http"
    "github.com/gorilla/websocket"
    "app/hub"
    "app/message"
)

var upgrader = websocket.Upgrader{}

func writeSocket(socket *websocket.Conn, c hub.Connection)  {
    defer socket.Close()
    for {
        m := <- c.Out
        socket.WriteJSON(&m)
    }
}

// Handler handles websocket connections at /ws
func Handler(w http.ResponseWriter, r *http.Request) {
    socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        panic(err)
	}
	defer socket.Close()

    c := hub.NewConnection()

    go writeSocket(socket, c)

    for {
        m := message.SocketMessage{}
        m.CreatedAt = time.Now().UTC()

		socket.ReadJSON(&m)

        if m.Action == "publish" {
            hub.Publish(m)
        }

        if m.Action == "subscribe" {
            hub.Subscribe(m.Event, c)
        }

        if m.Action == "unsubscribe" {
            hub.Unsubscribe(m.Event, c)
        }
	}
}
