package info

import (
	"github.com/exitcodezero/picloud/hub"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SocketHandler handles websocket connections at /info
func SocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer socket.Close()

	for now := range time.Tick(5 * time.Second) {
		im := hub.Manager.Info()
		im.CreatedAt = now.UTC()
		err := socket.WriteJSON(&im)
		if err != nil {
			break
		}
	}
}
