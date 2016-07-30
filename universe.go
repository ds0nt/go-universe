package main

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"

	"golang.org/x/net/websocket"

	"github.com/ds0nt/go-universe/socket"
	"github.com/ds0nt/go-universe/universe"
)

func main() {
	stop := make(chan struct{})
	//
	// err := termui.Init()
	// if err != nil {
	// 	panic(err)
	// }
	// defer termui.Close()
	//
	// mm := minimap.NewMiniMap()
	// handle key q
	// termui.Handle("/sys/kbd/q", func(termui.Event) {
	// 	stop <- struct{}{}
	// })

	// mPoints := []*minimap.MapPoint{}
	// for _, v := range u.Entities {
	// 	pos := v.Position()
	// 	mPoints = append(mPoints, &minimap.MapPoint{
	// 		X: int(pos.X),
	// 		Y: int(pos.Y),
	// 	})
	// }
	// mm.SetPoints(mPoints)
	// termui.Render(mm)

	u := universe.NewUniverse()

	p1 := newPlayer()
	p2 := newPlayer()
	p3 := newPlayer()
	p4 := newPlayer()
	p2.SetPosition(universe.Position{10, 10, 10})
	p3.SetPosition(universe.Position{5, 20, 50})
	p4.SetPosition(universe.Position{2, 11, 50})
	u.Add(p1, p2, p3, p4)

	// websocket interface
	go func() {
		log.Println("Starting Server")
		universeServer := socket.NewUniverseServer(u)
		http.Handle("/connect", websocket.Handler(universeServer.Server))
		err := http.ListenAndServe(":12345", nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	updateCh := make(chan *universe.UniverseTickParams)
	go u.Start(updateCh)
	defer u.Stop()

	ticker := time.NewTicker(time.Second / 60)
	lastTime := time.Now()
	for {
		select {
		case <-stop:
			return
		default:
			newTime := <-ticker.C
			updateCh <- &universe.UniverseTickParams{
				Time:  newTime,
				Delta: newTime.Sub(lastTime),
			}
			lastTime = newTime

		}
	}
}

// some objects
type Player struct {
	pos universe.Position
}

func newPlayer() *Player {
	return &Player{}
}

func (p *Player) Position() universe.Position {
	return p.pos
}

func (p *Player) SetPosition(p2 universe.Position) {
	p.pos.SetValue(p2)
}
