package hub

import (
    "app/message"
)


// Published is a channel where socket messages that need to be published to other
// connections are pushed
var Published chan message.SocketMessage

// Subscribed maps event strings to a list of channels that should receive the
// message when it is Published
var Subscribed map[string][]chan message.SocketMessage

func processSubscriptions(pub chan message.SocketMessage, sub map[string][]chan message.SocketMessage)  {
    for {
        message := <- pub
        for _, s := range sub[message.Event] {
            s <- message
        }
    }
}

func init()  {
    Published = make(chan message.SocketMessage)
    Subscribed = make(map[string][]chan message.SocketMessage)
    go processSubscriptions(Published, Subscribed)
}
