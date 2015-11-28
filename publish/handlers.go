package publish

import (
    "encoding/json"
	"github.com/exitcodezero/picloud/hub"
    "github.com/exitcodezero/picloud/message"
	"net/http"
	"time"
)

type publishBody struct {
    Event     string    `json:"event"`
	Data      string    `json:"data"`
}

// Handler publishes messages via HTTP
func Handler(w http.ResponseWriter, r *http.Request) {
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
        Action: "publish",
        Event: pb.Event,
        Data: pb.Data,
        CreatedAt: time.Now().UTC(),
    }

    hub.Manager.Publish(m)

    w.WriteHeader(http.StatusCreated)
}
