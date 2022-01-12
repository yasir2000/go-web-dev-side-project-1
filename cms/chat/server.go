package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
)

//WSUpgrader is used to upgrade the protocol to allow websockets

var WSUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Conn represents a websocket connection
type Conn struct {
	WS   *websocket.Conn
	Send chan string
}

// SendToHub method sends any message from our websocket connection to our hub
func (conn *Conn) SendToHub() {
	defer conn.WS.Close()
	for {
		_, msg, err := conn.WS.ReadMessage()
		if err != nil {
			// user has been dissconnected - probably just refreshed his browser
			// so just return
			return
		}
		DefaultHub.Echo <- string(msg)
	}
}

// RecieveFromHub sends message from our hub to our websocket connection
func (conn *Conn) RecieveFromHub() {
	defer conn.WS.Close()
	for {
		conn.Write(<-conn.Send)
	}
}

// Write method writes message
func (conn *Conn) Write(msg string) error {
	return conn.WS.WriteMessage(websocket.TextMessage, []byte(msg))
}

//WSHandler handles the HTTP req
func WSHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := WSUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	conn := &Conn{
		Send: make(chan string),
		WS:   ws,
	}
	DefaultHub.Join <- conn

	go conn.SendToHub()
	conn.RecieveFromHub()
}
