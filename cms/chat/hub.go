package chat

// DefaultHub is the default hub
var DefaultHub = NewHub()

// Hub is the data structure to be used to keep track of connections
type Hub struct {
	Join  chan *Conn
	Conns map[*Conn]bool
	Echo  chan string
}

// NewHub creates a new default hub
func NewHub() *Hub {
	return &Hub{
		Join:  make(chan *Conn),
		Conns: make(map[*Conn]bool),
		Echo:  make(chan string),
	}
}

// Start method starts our hub
func (hub *Hub) Start() {
	for {
		// Select here is a multiplexer for channels which will
		// wait for one of its cases to run
		select {
		case conn := <-hub.Join:
			DefaultHub.Conns[conn] = true
		case msg := <-hub.Echo:
			for conn := range hub.Conns {
				conn.Send <- msg
			}

		}
	}

}
