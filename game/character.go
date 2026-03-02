package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jizogames/horunpa/game/assets"
	"github.com/jizogames/horunpa/game/draw"
)

type Character struct {
	image *ebiten.Image
}

func (c *Character) Draw(screen *ebiten.Image) {
	draw.DrawAt(screen, c.image, 504, 144)
}

func NewCharacter() (*Character, error) {
	charaImage, err := draw.LoadImage(assets.Images, "images/horunpa.png")
	if err != nil {
		return nil, fmt.Errorf("キャラ画像のロードに失敗しました: %w", err)
	}

	c := &Character{
		image: charaImage,
	}

	return c, nil
}
