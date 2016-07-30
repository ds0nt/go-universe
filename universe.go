package main

import (
	"fmt"
	"log"
	"time"

	minimap "github.com/ds0nt/go-universe/minimap"
	termui "github.com/gizak/termui"
)

type UniverseTickParams struct {
	Time  time.Time
	Delta time.Duration
}

func main() {
	stop := make(chan struct{})
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()
	// build
	mm := minimap.NewMiniMap()
	// // build layout
	// termui.Body.AddRows(
	// 	termui.NewRow(),
	// 	termui.NewRow(
	// 		termui.NewCol(12, 0, mm)))

	// calculate layout
	// termui.Body.Align()

	// termui.Render(termui.Body)
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		stop <- struct{}{}
	})
	// handle key q
	u := newUniverse()

	p1 := newPlayer()
	p2 := newPlayer()
	p3 := newPlayer()
	p4 := newPlayer()
	p2.SetPosition(Position{10, 10, 10})
	p3.SetPosition(Position{5, 20, 50})
	p4.SetPosition(Position{2, 11, 50})
	u.Add(p1, p2, p3, p4)

	updateCh := make(chan *UniverseTickParams)
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
			updateCh <- &UniverseTickParams{
				Time:  newTime,
				Delta: newTime.Sub(lastTime),
			}
			lastTime = newTime

			mPoints := []*minimap.MapPoint{}
			for _, v := range u.Entities {
				pos := v.Position()
				mPoints = append(mPoints, &minimap.MapPoint{
					X: int(pos.X),
					Y: int(pos.Y),
				})
			}
			mm.SetPoints(mPoints)
			termui.Render(mm)
		}

	}

}

type Position struct {
	X float64
	Y float64
	Z float64
}

func (p *Position) String() string {
	return fmt.Sprintf("(%f, %f, %f)", p.X, p.Y, p.Z)
}

func (p *Position) SetValue(p2 Position) {
	p.X = p2.X
	p.Y = p2.Y
	p.Z = p2.Z
}

func newPosition() *Position {
	return &Position{}
}

type Entity interface {
	Position() Position
	SetPosition(p Position)
}

type Universe struct {
	Entities []Entity
	stop     chan struct{}
}

func newUniverse() *Universe {
	return &Universe{}
}

func (u *Universe) Add(e ...Entity) {
	u.Entities = append(u.Entities, e...)
}

func (u *Universe) Stop() {
	u.stop <- struct{}{}
}

const holeX = 30
const holeY = 30

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
			}
		}
	}
}

// some objects
type Player struct {
	pos Position
}

func newPlayer() *Player {
	return &Player{}
}

func (p *Player) Position() Position {
	return p.pos
}

func (p *Player) SetPosition(p2 Position) {
	p.pos.SetValue(p2)
}
