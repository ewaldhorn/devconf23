package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	red, green, blue uint8
}

func (g *Game) randomiseColours() {
	g.red = uint8(rand.Intn(255))
	g.green = uint8(rand.Intn(255))
	g.blue = uint8(rand.Intn(255))
}

func (g *Game) Update() error {
	g.randomiseColours()
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: g.red, G: g.green, B: g.blue, A: 0xff})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetTPS(1)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("1989 all over again")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
