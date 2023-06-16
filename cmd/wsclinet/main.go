package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
)

func main() {
	interrupt := make(chan os.Signal, 1)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1", Path: "/ws"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	// Go routine to read message sent from the WebSocket server.
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				return
			}
			fmt.Printf("received message: %s\n", message)
		}
	}()

	// Loop for sending message to WebSocket server until the connection is interrupted.
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, ""))
			if err != nil {
				log.Println("write close error:", err)
				return
			}
			select {
			case <-done:
			}
			return
		}
	}
}
