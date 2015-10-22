package socket

import (
    "github.com/gorilla/websocket"
    "app/message"
)

func writeSocket(c *websocket.Conn, out chan message.SocketMessage)  {
    defer c.Close()
    for {
        message := <- out
        err := c.WriteJSON(&message)
        if err != nil {
        	panic(err)
        }
    }
}
