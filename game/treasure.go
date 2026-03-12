package game

import (
	"fmt"
	"image"
	_ "image/png"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jizogames/horunpa/game/assets"
)

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

func LoadTreasureImages() error {
	filename := "treasure.png"

	f, err := assets.Images.Open(path.Join("images", filename))
	if err != nil {
		return fmt.Errorf("フォルダを開けませんでした: %w", err)
	}

	img, _, err := image.Decode(f)
	f.Close()
	if err != nil {
		return fmt.Errorf("画像のデコードに失敗しました: %w", err)
	}

	treasure := ebiten.NewImageFromImage(img)
	for y := 0; y < 6; y++ {
		tre := treasure.SubImage(image.Rect(0, y*26, 39, (y+1)*26)).(*ebiten.Image)
		treasureImg = append(treasureImg, tre)
	}

	return nil
}
