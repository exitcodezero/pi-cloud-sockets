package socket

import (
    "github.com/gorilla/websocket"
    "app/message"
)

func writeSocket(c *websocket.Conn, rec chan message.SocketMessage)  {
    defer c.Close()
    for {
        message := <- rec
        err := c.WriteJSON(&message)
        if err != nil {
        	panic(err)
        }
    }
}
