package socket

import (
	log "github.com/Sirupsen/logrus"

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

type BaseMessage struct {
	T string      `json:"t"`
	D interface{} `json:"d"`
}

type EnterMessageData struct {
	Username string
	Password string
}

func (u *UniverseServer) Server(ws *websocket.Conn) {
	msgChan := make(chan *BaseMessage)
	defer close(msgChan)

	go func() {
		for {
			var baseMsg BaseMessage
			err := websocket.JSON.Receive(ws, &baseMsg)
			if err != nil {
				log.Println("Error", err)
			}
			switch baseMsg.T {
			case "enter":
				m, ok := baseMsg.D.(EnterMessageData)
				if !ok {
					log.Println("Bad EnterMessageData")
				}
				u.HandleEnter(ws, &m, msgChan)
			}
		}
	}()

	for {
		outMsg := <-msgChan
		websocket.JSON.Send(ws, outMsg)
	}
}

func (u *UniverseServer) HandleEnter(ws *websocket.Conn, in *EnterMessageData, out chan<- *BaseMessage) {

}
