package socket

import (
	"github.com/ds0nt/go-universe/universe"

	"golang.org/x/net/websocket"
)

type UniverseServer struct {
	universe *universe.Universe
}

func NewUniverseServer(universe *universe.Universe) *UniverseServer {
	u := &UniverseServer{
		universe: universe,
	}
	return u
}

func (u *UniverseServer) Server(ws *websocket.Conn) {
	ws.Read()
}
