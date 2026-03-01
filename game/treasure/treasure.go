package treasure

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jizogames/horunpa/game/ui"
)

type Manager struct {
	windows []*ui.Window
}

func (m *Manager) Draw(screen *ebiten.Image) {
	for _, win := range m.windows {
		win.Draw(screen)
	}
}

func NewManager() (*Manager, error) {
	windowAreaStartX := 9
	windowAreaStartY := 9
	windowWidth := 117
	windowHeight := 78
	windowGap := 9
	numberOfWindows := 3

	var windows []*ui.Window
	for i := 0; i < numberOfWindows; i++ {
		y := windowAreaStartY + i*(windowHeight+windowGap)
		rect := ui.NewRect(windowAreaStartX, y, windowWidth, windowHeight)
		frameColor := color.RGBA{118, 92, 71, 255}
		window := ui.NewWindow(rect, frameColor, nil, 4)
		window.Activate()

		windows = append(windows, window)
	}

	m := &Manager{
		windows: windows,
	}

	return m, nil
}
