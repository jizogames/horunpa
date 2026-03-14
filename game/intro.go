package game

import "github.com/hajimehoshi/ebiten/v2"

type Task struct {
	frame     int
	Triggered bool
	Action    func(i *Intro)
}

type Intro struct{}

func (i *Intro) Update() error {
	return nil
}

func (i *Intro) Draw(screen *ebiten.Image) {}
