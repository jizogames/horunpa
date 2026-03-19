package game

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"math/rand"
	"path"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jizogames/horunpa/game/assets"
	"github.com/jizogames/horunpa/game/audio"
	"github.com/jizogames/horunpa/game/draw"
)

type Game struct {
	scene Scene
}

func (g *Game) Update() error {
	if g.scene == nil {
		g.scene = NewIntro()
	}

	switch g.scene.Msg() {
	case GAMESTATE_MSG_REQ_TITLE:
		g.scene = NewTitle()
	case GAMESTATE_MSG_REQ_MAIN:
		g.scene = NewGameScene()
	}

	g.scene.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
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

type GameStateMsg int

const (
	GAMESTATE_MSG_NONE GameStateMsg = iota
	GAMESTATE_MSG_REQ_INTRO
	GAMESTATE_MSG_REQ_TITLE
	GAMESTATE_MSG_REQ_MAIN
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
	Msg() GameStateMsg
}

type GameScene struct {
	audio *audio.Manager

	wall      *Wall
	chara     *Character
	treasures []*Treasure

	gameStateMsg GameStateMsg
}

func (g *GameScene) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cx, cy := ebiten.CursorPosition()
		if cx < 135 || cy < 9 || cx > 495 || cy > 261 {
			return
		}

		x := (cx - 135) / 36
		y := (cy - 9) / 36

		if g.wall.Cells[y][x].HP > 0 {
			g.wall.Cells[y][x].HP--
		}
	}
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{224, 235, 175, 255})
	g.wall.Draw(screen)
	g.chara.Draw(screen)
}

func (g *GameScene) Msg() GameStateMsg {
	return g.gameStateMsg
}

func NewGameScene() *GameScene {
	const audioSampleRate int = 48000
	audioManager, err := audio.NewManager(audioSampleRate)
	if err != nil {
		panic(err)
	}

	chara, err := NewCharacter()
	if err != nil {
		panic(err)
	}

	LoadCellImages()
	wall, err := NewWall()
	if err != nil {
		panic(err)
	}

	LoadTreasureImages()
	treasures := make([]*Treasure, 3)
	for i := 0; i < 3; i++ {
		id := rand.Intn(3)
		treasures[i] = &Treasure{
			ID: id,
		}
	}

	g := &GameScene{
		audio: audioManager,

		wall:      wall,
		chara:     chara,
		treasures: treasures,
	}

	if err := g.audio.PlayBGM("bgm"); err != nil {
		panic(err)
	}

	return g
}

type Title struct {
	logo         *ebiten.Image
	gameStateMsg GameStateMsg
}

func (t *Title) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		t.gameStateMsg = GAMESTATE_MSG_REQ_MAIN
	}
}

func (t *Title) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{224, 235, 175, 255})
	draw.DrawAt(screen, t.logo, 150, 90)
}

func (t *Title) Msg() GameStateMsg {
	return t.gameStateMsg
}

func NewTitle() *Title {
	logo, err := draw.LoadImage(assets.Images, "images/title.png")
	if err != nil {
		panic(err)
	}

	t := &Title{
		logo: logo,
	}
	return t
}
