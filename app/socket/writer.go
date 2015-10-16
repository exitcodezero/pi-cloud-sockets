package socket

import (
    "log"
    "github.com/gorilla/websocket"
    "app/message"
)

func writeSocket(c *websocket.Conn, pub chan message.SocketMessage)  {
    for {
        message := <- pub
        log.Println(message)
        err := c.WriteJSON(&message)
        if err != nil {
        	panic(err)
        }
    }
}
