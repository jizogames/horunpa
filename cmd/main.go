package main

import "github.com/jizogames/horunpa/game"

func main() {
	if err := game.Run(); err != nil {
		panic(err)
	}
}
