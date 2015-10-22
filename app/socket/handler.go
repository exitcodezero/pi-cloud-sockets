package socket

import (
    "time"
    "net/http"
    "github.com/gorilla/websocket"
    "app/hub"
    "app/message"
)

var upgrader = websocket.Upgrader{}

// Handler handles websocket connections at /ws
func Handler(w http.ResponseWriter, r *http.Request) {
    socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        panic(err)
	}
	defer socket.Close()

    out := make(chan message.SocketMessage)

    go writeSocket(socket, out)

    for {
        m := message.SocketMessage{}
        m.CreatedAt = time.Now().UTC()

		socket.ReadJSON(&m)

        if m.Action == "publish" {
            hub.Published <- m
        }

        if m.Action == "subscribe" {
            hub.Subscribed[m.Event] = append(hub.Subscribed[m.Event], out)
        }
	}

}
