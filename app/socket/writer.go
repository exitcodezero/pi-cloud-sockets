package socket

import (
    "log"
    "github.com/gorilla/websocket"
    "app/message"
)

func writeSocket(c *websocket.Conn, rec chan message.SocketMessage)  {
    defer c.Close()
    for {
        message := <- rec
        log.Println(message)
        err := c.WriteJSON(&message)
        if err != nil {
        	panic(err)
        }
    }
}
