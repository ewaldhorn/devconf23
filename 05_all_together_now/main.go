package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

var titleImage *ebiten.Image
var gopherImage *ebiten.Image

func init() {
	var err error
	titleImage, _, err = ebitenutil.NewImageFromFile("assets/devconfTitle.png")
	if err != nil {
		log.Fatal(err)
	}

	gopherImage, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	gopherX, gopherY         int
	gopherXIncr, gopherYIncr int
	gopherAngle              float64
}

func (g *Game) Update() error {
	g.updateGopherPosition()
	return nil
}

func (g *Game) updateGopherPosition() {
	if g.gopherX > 800 || g.gopherX < 5 {
		g.gopherXIncr *= -1
		g.gopherAngle = 0.0
	}

	g.gopherX += g.gopherXIncr

	if g.gopherY > 600 || g.gopherY < 5 {
		g.gopherYIncr *= -1
	}

	g.gopherX += g.gopherXIncr
	g.gopherY += g.gopherYIncr

	g.gopherAngle += 0.001
}

func (g *Game) Draw(screen *ebiten.Image) {
	titleImageDrawingOptions := &ebiten.DrawImageOptions{}
	titleImageDrawingOptions.GeoM.Scale(0.75, 1.0)
	screen.DrawImage(titleImage, titleImageDrawingOptions)

	gopherDrawingOptions := &ebiten.DrawImageOptions{}
	gopherDrawingOptions.GeoM.Translate(float64(g.gopherX), float64(g.gopherY))
	gopherDrawingOptions.GeoM.Scale(0.5, 0.5)
	gopherDrawingOptions.GeoM.Rotate(g.gopherAngle)
	screen.DrawImage(gopherImage, gopherDrawingOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func initGame() *Game {
	game := Game{}

	game.gopherXIncr = 1
	game.gopherYIncr = 1

	game.gopherX = 30
	game.gopherY = 30

	return &game
}

func main() {
	ebiten.SetTPS(25)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Whoo Whoo!")

	if err := ebiten.RunGame(initGame()); err != nil {
		log.Fatal(err)
	}
}
