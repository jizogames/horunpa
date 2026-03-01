package assets

import "embed"

var (
	//go:embed icons/*
	Icons embed.FS

	//go:embed fonts/*
	Fonts embed.FS
)
