package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"math"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for radius := 10; radius < 70; radius += 2 {
		g.drawCircle(screen, 160, 100, radius, color.RGBA{R: 128, G: 255, B: 128, A: 255})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 200
}

// Shamelessly stolen from https://golangprojectstructure.com/creating-cool-games-with-ebiten-in-go/
func (g *Game) drawCircle(screen *ebiten.Image, x, y, radius int, clr color.Color) {
	radius64 := float64(radius)
	minAngle := math.Acos(1 - 1/radius64)
	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius64 * math.Cos(angle)
		yDelta := radius64 * math.Sin(angle)
		x1 := int(math.Round(float64(x) + xDelta))
		y1 := int(math.Round(float64(y) + yDelta))
		screen.Set(x1, y1, clr)
	}
}

func main() {
	ebiten.SetWindowSize(320, 200)
	ebiten.SetWindowTitle("Back to Basics")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
