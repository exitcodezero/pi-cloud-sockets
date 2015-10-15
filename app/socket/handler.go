package socket

import (
    "net/http"
    "github.com/gorilla/websocket"
    "app/rabbit"
)

type socketMessage struct {
    Action string `json:"action"`
    Event string `json:"event"`
    Data string `json:"data"`
}

var upgrader = websocket.Upgrader{}

// Handler handles websocket connections at /ws
func Handler(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        panic(err)
	}
	defer c.Close()

    ch, err := rabbit.Conn.Channel()
    if err != nil {
        panic(err)
    }
    defer ch.Close()

    for {
        message := socketMessage{}

		err := c.ReadJSON(&message)
		if err != nil {
			panic(err)
		}

		err = c.WriteJSON(&message)
		if err != nil {
			panic(err)
		}
	}
}
