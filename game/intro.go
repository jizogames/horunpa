package game

import "github.com/hajimehoshi/ebiten/v2"

type Intro struct {
	frame int

	tasks   []Task
	sprites []Sprite
}

func (i *Intro) Update() {
	i.frame++

	for idx := range i.tasks {
		task := &i.tasks[idx]
		if !task.Triggered && i.frame >= task.Frame {
			task.Action(i)
			task.Triggered = true
		}
	}

	dst := i.sprites[:0]
	for _, s := range i.sprites {
		s.Update()
		if s.Alive() {
			dst = append(dst, s)
		}
	}
	i.sprites = dst
}

func (i *Intro) Draw(screen *ebiten.Image) {
	for _, s := range i.sprites {
		s.Draw(screen)
	}
}

// タスク.
type Task struct {
	Frame     int
	Triggered bool
	Action    func(i *Intro)
}

// スプライト.
type Sprite interface {
	Update()
	Draw(screen *ebiten.Image)
	Alive() bool
}
