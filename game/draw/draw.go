package draw

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(fs embed.FS, filename string) (*ebiten.Image, error) {
	f, err := fs.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("フォルダを開けませんでした: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("画像のデコードに失敗しました(%s): %w", filename, err)
	}

	return ebiten.NewImageFromImage(img), nil
}

func DrawAt(dst *ebiten.Image, src *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	dst.DrawImage(src, op)
}
