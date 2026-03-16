package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jizogames/horunpa/game/assets"
	"github.com/jizogames/horunpa/game/draw"
)

type Intro struct {
	frame        int
	gameStateMsg GameStateMsg

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

func (i *Intro) Msg() GameStateMsg {
	return i.gameStateMsg
}

func (i *Intro) AllTasksFinished() {
	i.gameStateMsg = GAMESTATE_MSG_REQ_TITLE
}

func NewIntro() *Intro {
	i := &Intro{}

	i.tasks = []Task{
		{
			Frame: 0,
			Action: func(i *Intro) {
				i.sprites = append(i.sprites, NewLogoSprite(35, 50, 240))
			},
		},
		{
			Frame: 200,
			Action: func(i *Intro) {
				i.AllTasksFinished()
			},
		},
	}

	return i
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

// スプライト：ロゴ.
type LogoSprite struct {
	img   *ebiten.Image
	x, y  int
	alpha float32

	life  int
	alive bool
}

func (l *LogoSprite) Update() {
	l.life--
	if l.life <= 0 {
		l.alive = false
		return
	}

	if l.life >= 100 {
		l.alpha += 0.02
		if l.alpha > 1 {
			l.alpha = 1
		}
	} else if l.life <= 50 {
		l.alpha -= 0.02
		if l.alpha < 0 {
			l.alpha = 0
		}
	}
}

func (l *LogoSprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(l.x), float64(l.y))
	op.ColorScale.ScaleAlpha(l.alpha)
	screen.DrawImage(l.img, op)
}

func (l *LogoSprite) Alive() bool {
	return l.alive
}

func NewLogoSprite(x, y, life int) *LogoSprite {
	img, err := draw.LoadImage(assets.Images, "images/logo.png")
	if err != nil {
		panic(err)
	}

	l := &LogoSprite{
		img: img,

		x:     x,
		y:     y,
		life:  life,
		alive: true,
	}

	return l
}
