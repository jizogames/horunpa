package assets

import "embed"

var (
	//go:embed icons/*
	Icons embed.FS

	//go:embed fonts/*
	Fonts embed.FS

	//go:embed audio/*
	Audio embed.FS

	//go:embed images/*
	Images embed.FS
)
