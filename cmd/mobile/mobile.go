package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/jizogames/horunpa/game"
)

func init() {
	game, err := game.NewGame()
	if err != nil {
		panic(err)
	}
	mobile.SetGame(game)
}

func Dummy() {}
