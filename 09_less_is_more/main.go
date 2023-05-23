package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
)

const (
	screenWidth  = 400
	screenHeight = 400
	bouncerSize  = 4
	bouncerCount = 30
)

type Bouncer struct {
	positionX, positionY float64
	movementX, movementY float64
	colour               color.RGBA
}

func (b *Bouncer) init() {
	b.positionX = float64(rand.Intn(screenWidth))
	b.positionY = float64(rand.Intn(screenWidth))

	if rand.Int()%2 == 0 {
		b.movementX = 1
		b.movementY = -1
	} else {
		b.movementX = -1
		b.movementY = 1
	}

	b.colour = color.RGBA{
		R: uint8(rand.Intn(255)),
		G: uint8(rand.Intn(255)),
		B: uint8(rand.Intn(255)),
		A: 255,
	}
}

func (b *Bouncer) update() {
	b.positionX += b.movementX
	b.positionY += b.movementY

	if b.positionX >= screenWidth-bouncerSize || b.positionX <= bouncerSize {
		b.movementX *= -1
	}
	if b.positionY >= screenHeight-bouncerSize || b.positionY <= bouncerSize {
		b.movementY *= -1
	}
}

type Game struct {
	bouncers []Bouncer
}

func (g *Game) initBouncers() {
	g.bouncers = make([]Bouncer, bouncerCount)

	for bouncerPosition := range g.bouncers {
		tmpBouncer := Bouncer{}
		tmpBouncer.init()
		g.bouncers[bouncerPosition] = tmpBouncer
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	for pos, bouncer := range g.bouncers {
		bouncer.update()
		g.bouncers[pos] = bouncer
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, bouncer := range g.bouncers {
		x := float32(bouncer.positionX)
		y := float32(bouncer.positionY)
		vector.DrawFilledCircle(screen, x, y, bouncerSize, bouncer.colour, true)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Less effort is good, right?")

	game := Game{}
	game.initBouncers()

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
