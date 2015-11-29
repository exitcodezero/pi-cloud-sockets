package publish

import (
	"encoding/json"
	"github.com/exitcodezero/picloud/hub"
	"github.com/exitcodezero/picloud/message"
	"github.com/gorilla/context"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type publishBody struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

// HandlerHTTP publishes messages via HTTP
func HandlerHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	pb := publishBody{}
	err := decoder.Decode(&pb)
	if err != nil {
		panic(err)
	}

	if pb.Event == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	m := message.SocketMessage{
		Action:    "publish",
		Event:     pb.Event,
		Data:      pb.Data,
		CreatedAt: time.Now().UTC(),
	}

	hub.Manager.Publish(m)

	w.WriteHeader(http.StatusCreated)
}


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandlerSocket handles inbound websocket connections only at /publish
func HandlerSocket(w http.ResponseWriter, r *http.Request) {
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

	// Handle inbound publish messages
	for {
		m := message.SocketMessage{
			Action: "publish",
			CreatedAt: time.Now().UTC(),
		}

		err = socket.ReadJSON(&m)
		if err != nil {
			break
		}

		hub.Manager.Publish(m)
	}
}
