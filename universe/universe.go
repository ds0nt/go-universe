package universe

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
)

type Entity interface {
	Position() Position
	SetPosition(p Position)
}

type Universe struct {
	Entities []Entity
	stop     chan struct{}
}

func NewUniverse() *Universe {
	return &Universe{}
}

func (u *Universe) Add(e ...Entity) {
	u.Entities = append(u.Entities, e...)
}

func (u *Universe) Stop() {
	u.stop <- struct{}{}
}

type UniverseTickParams struct {
	Time  time.Time
	Delta time.Duration
}

func (u *Universe) Start(update <-chan *UniverseTickParams) {
	for {
		select {
		case <-u.stop:
			log.Println("Stopping the Universe...")
			return
		default:
			frame := <-update
			for _, v := range u.Entities {
				p := v.Position()
				p.X += 1 * frame.Delta.Seconds()
				v.SetPosition(p)
				fmt.Print("˰")
			}
		}
	}
}
