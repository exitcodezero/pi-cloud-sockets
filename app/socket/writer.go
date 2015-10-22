package socket

import (
    "github.com/gorilla/websocket"
    "app/message"
)

func writeSocket(socket *websocket.Conn, out chan message.SocketMessage)  {
    defer socket.Close()
    for {
        m := <- out
        socket.WriteJSON(&m)
    }
}
