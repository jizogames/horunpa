package audio

import (
	"embed"
	"fmt"
	"path"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/jizogames/horunpa/game/assets"
)

type seEntry struct {
	data []byte
	pool []*audio.Player
}

type Manager struct {
	ctx *audio.Context

	bgm map[string]*audio.Player
	se  map[string]*seEntry

	mute      bool
	bgmVolume float64
	seVolume  float64

	currentBGM string
}

func (m *Manager) Load(fs embed.FS, dir string) error {
	ents, err := fs.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("フォルダを開けませんでした: %w", err)
	}

	for _, ent := range ents {
		name := ent.Name()
		ext := filepath.Ext(name)
		if ext != ".mp3" && ext != ".wav" {
			continue
		}

		f, err := fs.Open(path.Join(dir, name))
		if err != nil {
			return fmt.Errorf("ファイルを開けませんでした: %w", err)
		}
		key := name[:len(name)-len(ext)]

		switch ext {
		case ".mp3":
			stream, err := mp3.DecodeWithSampleRate(m.ctx.SampleRate(), f)
			f.Close()
			if err != nil {
				return fmt.Errorf("MP3のデコードに失敗しました(%s): %w", name, err)
			}

			loop := audio.NewInfiniteLoop(stream, stream.Length())
			p, err := m.ctx.NewPlayer(loop)
			if err != nil {
				return fmt.Errorf("BGMプレイヤーの初期化に失敗しました: %w", err)
			}
			m.bgm[key] = p
		case ".wav":
		}
	}
	return nil
}

func (m *Manager) Close() error {
	for _, p := range m.bgm {
		if err := p.Close(); err != nil {
			return fmt.Errorf("BGMプレイヤーを閉じることに失敗しました: %w", err)
		}
	}

	for _, e := range m.se {
		for _, p := range e.pool {
			if err := p.Close(); err != nil {
				return fmt.Errorf("SEプレイヤーを閉じることに失敗しました: %w", err)
			}
		}
	}
	return nil
}

func (m *Manager) SetBGMVolume(volume float64) {
	m.bgmVolume = volume
	if m.mute {
		return
	}

	if p, ok := m.bgm[m.currentBGM]; ok && p.IsPlaying() {
		p.SetVolume(volume)
	}
}

func (m *Manager) SetSEVolume(volume float64) {
	m.seVolume = volume
}

func (m *Manager) PauseBGM() {
	for _, p := range m.bgm {
		p.Pause()
	}
}

func (m *Manager) PlayBGM(name string) error {
	if m.mute {
		return nil
	}

	p, ok := m.bgm[name]
	if !ok {
		return fmt.Errorf("BGMが見つかりませんでした(%s)", name)
	}

	m.PauseBGM()
	m.currentBGM = name

	p.SetVolume(m.bgmVolume)
	if err := p.Rewind(); err != nil {
		return fmt.Errorf("BGMの巻き戻しに失敗しました: %w", err)
	}

	p.Play()
	return nil
}

func (m *Manager) PlaySE(name string) {
	if m.mute {
		return
	}

	e, ok := m.se[name]
	if !ok {
		fmt.Printf("SEが見つかりませんでした(%s)", name)
	}

	var p *audio.Player
	for _, cand := range e.pool {
		if !cand.IsPlaying() {
			p = cand
			break
		}
	}

	if p == nil {
		np := m.ctx.NewPlayerFromBytes(e.data)
		e.pool = append(e.pool, np)
		p = np
	}

	p.SetVolume(m.seVolume)
	p.Rewind()
	p.Play()
}

func (m *Manager) Mute() {
	m.mute = true
	m.PauseBGM()
}

func NewManager(sampleRate int) (*Manager, error) {
	m := &Manager{
		ctx:       audio.NewContext(sampleRate),
		bgm:       map[string]*audio.Player{},
		se:        map[string]*seEntry{},
		bgmVolume: 1.0,
		seVolume:  1.0,
	}

	if err := m.Load(assets.Audio, "audio"); err != nil {
		return nil, fmt.Errorf("オーディオの読み込みに失敗しました: %w", err)
	}

	return m, nil
}
