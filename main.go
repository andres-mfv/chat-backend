package main

import (
	"flag"
	"github.com/andres-mfv/chat-backend/src/room"
	"github.com/andres-mfv/chat-backend/src/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()

	roomService := service.NewRoomService()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c)
	})
	r.POST("/join", func(c *gin.Context) {

	})

	r.POST("/rooms/create", func(c *gin.Context) {
		req := room.Request{}
		if err := c.BindJSON(&req); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		err := roomService.CreateRoom(c, &req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
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
