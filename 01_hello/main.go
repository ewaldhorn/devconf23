package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// See https://ebitengine.org/en/tour/hello_world.html for the original.
// I took some slight liberties with the code for this talk.

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello Cape Town!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100, 50
}

func main() {
	ebiten.SetWindowSize(320, 200)
	ebiten.SetWindowTitle("DevConf 2023")

	err := ebiten.RunGame(&Game{})

	if err != nil {
		log.Fatal(err)
	}
}
