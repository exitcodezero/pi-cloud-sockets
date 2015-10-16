package socket

import (
    "net/http"
    "github.com/gorilla/websocket"
    "app/hub"
)

var upgrader = websocket.Upgrader{}

// Handler handles websocket connections at /ws
func Handler(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        panic(err)
	}
	defer c.Close()

    go writeSocket(c, hub.Published)

    readSocket(c, hub.Published)

}
