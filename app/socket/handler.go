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
        message := message.SocketMessage{}
        message.CreatedAt = time.Now().UTC()

		err := c.ReadJSON(&message)
		if err != nil {
			panic(err)
		}

        if message.Action == "publish" {
            hub.Published <- message
        }

        if message.Action == "subscribe" {
            hub.Subscribed[message.Event] = append(hub.Subscribed[message.Event], received)
        }
	}

}
