package hub

import (
    "app/message"
)

var publish chan message.SocketMessage

var subscribed map[string][]Connection

func processSubscriptions(pub chan message.SocketMessage, sub map[string][]Connection)  {
    for {
        message := <- pub
        for _, c := range sub[message.Event] {
            c.Out <- message
        }
    }
}

func init() {
    publish = make(chan message.SocketMessage)
    subscribed = make(map[string][]Connection)
    for i := 0; i < 5; i++ {
        go processSubscriptions(publish, subscribed)
    }
}

// Publish adds a SocketMessage to the Publish channel
func Publish(m message.SocketMessage) {
    publish <- m
}

// Subscribe adds a Connection to an array for the event key
func Subscribe(event string, c Connection)  {
    existing := findConnectionIndex(c, subscribed[event])
    if existing == -1 {
        subscribed[event] = append(subscribed[event], c)
    }
}

// Unsubscribe removes a Connection from the array for the event key
func Unsubscribe(event string, c Connection) {
    findConnectionIndex(c, subscribed[event])
}
