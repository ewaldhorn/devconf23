package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 100; x < 200; x++ {
		for y := 100; y < 200; y++ {
			col := color.RGBA{R: uint8(x), G: 0, B: uint8(y), A: 255}
			screen.Set(x, y, col)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("The what now?")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
