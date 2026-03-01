package game

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"path"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jizogames/horunpa/game/assets"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{224, 235, 175, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 630, 270
}

func SetIcons(fs embed.FS, dir string) error {
	ents, err := fs.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("フォルダを読み込めませんでした: %w", err)
	}

	var icons []image.Image
	for _, ent := range ents {
		name := ent.Name()
		ext := filepath.Ext(name)
		if ext != ".png" {
			continue
		}

		f, err := fs.Open(path.Join(dir, name))
		if err != nil {
			return fmt.Errorf("ファイルを開けませんでした(%s): %w", name, err)
		}

		img, _, err := image.Decode(f)
		f.Close()
		if err != nil {
			return fmt.Errorf("画像のデコードに失敗しました(%s): %w", name, err)
		}
		icons = append(icons, img)
	}
	ebiten.SetWindowIcon(icons)
	return nil
}

func NewGame() (*Game, error) {
	g := &Game{}

	ebiten.SetWindowTitle("ほるんぱ")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := SetIcons(assets.Icons, "icons"); err != nil {
		return nil, fmt.Errorf("アイコンの設定に失敗しました: %w", err)
	}

	return g, nil
}

func Run() error {
	g, err := NewGame()
	if err != nil {
		return fmt.Errorf("ゲームの初期化に失敗しました: %w", err)
	}

	if err := ebiten.RunGame(g); err != nil {
		return fmt.Errorf("ゲームの実行に失敗しました: %w", err)
	}

	return nil
}
