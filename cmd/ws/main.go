package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type client struct {
	conn *net.Conn
	send chan []byte
}

type hub struct {
	clients    map[*client]bool
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

func (h *hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func handleTCPRequest(conn net.Conn, h *hub) {
	defer conn.Close()

	bufReader := bufio.NewReader(conn)
	for {
		message, err := bufReader.ReadBytes('\n')
		if err != nil {
			log.Println("Error reading:", err.Error())
			return
		}

		log.Printf("Received message from TCP client: %s", message)

		h.broadcast <- message
	}
}

func handleWebSocketRequest(w http.ResponseWriter, r *http.Request, h *hub) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &client{conn: conn, send: make(chan []byte)}

	h.register <- client

	go func() {
		defer conn.Close()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				h.unregister <- client
				log.Println(err)
				break
			}

			log.Printf("Received message from WebSocket client: %s", message)

			h.broadcast <- message
		}
	}()

	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				return
			}
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func main() {
	h := hub{
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
	go h.run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocketRequest(w, r, &h)
	})

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	defer l.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			continue
		}

		go handleTCPRequest(conn, &h)
	}
}
