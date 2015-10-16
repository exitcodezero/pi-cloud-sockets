package socket

import (
    "log"
    "github.com/gorilla/websocket"
    "app/message"
)

func readSocket(c *websocket.Conn, pub chan message.SocketMessage)  {

    for {
        message := message.SocketMessage{}

		err := c.ReadJSON(&message)
		if err != nil {
			panic(err)
		}

        if message.Action == "publish" {
            log.Println(message.Action)
            pub <- message
        }
	}

}
