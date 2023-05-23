package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
	"math/rand"
)

const (
	screenWidth  = 400
	screenHeight = 400
	bouncerSize  = 4
	bouncerCount = 50
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

func (g *Game) Draw(screen *ebiten.Image) {
	for _, bouncer := range g.bouncers {
		x := int(math.Round(bouncer.positionX))
		y := int(math.Round(bouncer.positionY))
		g.drawCircle(screen, x, y, bouncerSize, bouncer.colour)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Round and round we go!")

	game := Game{}
	game.initBouncers()

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
