package game

import (
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jizogames/horunpa/game/assets"
	"github.com/jizogames/horunpa/game/draw"
)

var (
	CellImages []*ebiten.Image
)

const (
	WallWidth  = 10
	WallHeight = 7
)

type Cell struct {
	HP int

	TreasureIndex int // -1 なら何もない.
	PieceX        int
	PieceY        int
}

type Wall struct {
	Width, Height int

	Cells [][]Cell
}

func (w *Wall) Draw(screen *ebiten.Image) {
	wallX := 135
	wallY := 9
	for y := 0; y < WallHeight; y++ {
		row := w.Cells[y]
		for x := 0; x < WallWidth; x++ {
			cellHP := row[x].HP

			px := x*36 + wallX
			py := y*36 + wallY
			draw.DrawAt(screen, CellImages[cellHP], px, py)
		}
	}
}

func NewWall() (*Wall, error) {
	cells := make([][]Cell, WallHeight)
	for y := 0; y < WallHeight; y++ {
		row := make([]Cell, WallWidth)
		for x := 0; x < WallWidth; x++ {
			hp := rand.Intn(3) + 1
			row[x] = Cell{
				HP: hp,
			}
		}
		cells[y] = row
	}

	w := &Wall{
		Cells: cells,
	}
	return w, nil
}

func LoadCellImages() error {
	filename := "cell.png"

	f, err := assets.Images.Open(path.Join("images", filename))
	if err != nil {
		return fmt.Errorf("フォルダを開けませんでした: %w", err)
	}

	img, _, err := image.Decode(f)
	f.Close()
	if err != nil {
		return fmt.Errorf("画像のデコードに失敗しました: %w", err)
	}

	cellImage := ebiten.NewImageFromImage(img)
	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			cell := cellImage.SubImage(image.Rect(x*36, y*36, (x+1)*36, (y+1)*36)).(*ebiten.Image)
			CellImages = append(CellImages, cell)
		}
	}

	return nil
}
