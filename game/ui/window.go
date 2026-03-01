package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jizogames/horunpa/game/draw"
)

type Rect struct {
	x, y          int
	width, height int
}

func NewRect(x, y, w, h int) Rect {
	return Rect{
		x:      x,
		y:      y,
		width:  w,
		height: h,
	}
}

type Window struct {
	rect Rect

	frame      *ebiten.Image
	inner      *ebiten.Image
	frameWidth int

	isActive bool
}

func (w *Window) Draw(screen *ebiten.Image) {
	if !w.isActive {
		return
	}

	draw.DrawAt(screen, w.frame, w.rect.x, w.rect.y)
	if w.inner != nil {
		draw.DrawAt(screen, w.inner, w.rect.x+w.frameWidth, w.rect.y+w.frameWidth)
	}
}

func (w *Window) SetInnerImage(image *ebiten.Image) {
	w.inner = image
}

func (w *Window) Activate() {
	w.isActive = true
}

func (w *Window) Inactivate() {
	w.isActive = false
}

func NewWindow(rect Rect, frameColor color.RGBA, inner *ebiten.Image, frameWidth int) *Window {
	f := ebiten.NewImage(rect.width, rect.height)
	f.Fill(frameColor)

	w := &Window{
		rect: rect,

		frame:      f,
		inner:      inner,
		frameWidth: frameWidth,
	}

	return w
}
