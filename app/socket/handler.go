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
    c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        panic(err)
	}
	defer c.Close()

    // each socket connection has a 'received' channel
    received := make(chan message.SocketMessage)

    // all messages pushed to the 'received' channel
    // are written out to the socket
    go writeSocket(c, received)

    // read incoming messages from the socket
    for {
        m := message.SocketMessage{}
        m.CreatedAt = time.Now().UTC()

		err := c.ReadJSON(&m)
		if err != nil {
			panic(err)
		}

        if m.Action == "publish" {
            hub.Published <- m
        }

        if m.Action == "subscribe" {
            hub.Subscribed[m.Event] = append(hub.Subscribed[m.Event], received)
        }
	}

}
