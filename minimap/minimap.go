package minimap

import (
	"sync"

	"github.com/gizak/termui"
)

type MiniMap struct {
	sync.Mutex
	termui.Block
	PlayerColor     termui.Attribute
	PlayerRune      rune
	BackgroundColor termui.Attribute
	Scale           float64
	Points          []*MapPoint
}

type MapPoint struct {
	X int
	Y int
}

// NewMiniMap returns a new *MiniMap with current theme.
func NewMiniMap() *MiniMap {
	mm := &MiniMap{Block: *termui.NewBlock()}
	mm.PlayerColor = termui.ThemeAttr("barchart.bar.bg")
	mm.PlayerRune = 'x'
	mm.BackgroundColor = termui.ThemeAttr("barchart.bar.bg")
	mm.Height = 100
	mm.Width = 100
	mm.Scale = 1
	return mm
}

func (mm *MiniMap) SetPoints(points []*MapPoint) {
	mm.Lock()
	mm.Points = points
	mm.Unlock()
}

// Buffer implements Bufferer interface.
func (mm *MiniMap) Buffer() termui.Buffer {
	buf := mm.Block.Buffer()
	c := termui.Cell{
		Ch: mm.PlayerRune,
		Fg: mm.PlayerColor,
		Bg: mm.BackgroundColor,
	}

	mm.Lock()
	for _, p := range mm.Points {
		buf.Set(p.X, p.Y, c)
	}
	mm.Unlock()

	return buf
}
