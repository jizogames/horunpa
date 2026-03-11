package game

import "github.com/hajimehoshi/ebiten/v2"

var (
	treasureImg []*ebiten.Image
)

type Position struct {
	X, Y int
}

type Treasure struct {
	Position

	ID int
}

func (t *Treasure) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.X), float64(t.Y))
	screen.DrawImage(treasureImg[t.ID], op)
}
