package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type Server struct {
	peers map[*websocket.Conn]bool
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, World")
	w.Write([]byte{})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Couldn't upgrade connection %s", err)
		return
	}

	log.Printf("New connection %s", conn.RemoteAddr())
	server.peers[conn] = true

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Couldn't read message")
			break
		}

		log.Printf("Received: %s", message)
		server.broadcast(message)

	}
}

func NewServer() *Server {
	return &Server{
		peers: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) broadcast(msg []byte) {

	for conn := range s.peers {
		go func(conn *websocket.Conn) {
			conn.WriteMessage(websocket.TextMessage, msg)
		}(conn)
	}
}

var server = NewServer()

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/socket", WsHandler)
	http.ListenAndServe("localhost:5000", nil)
}
