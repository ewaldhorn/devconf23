package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
)

const (
	screenWidth  = 400
	screenHeight = 400
	bouncerCount = 30
)

var (
	lineWidth float32 = 2.0
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

	if b.positionX >= screenWidth || b.positionX <= 1 {
		b.movementX *= -1
	}
	if b.positionY >= screenHeight || b.positionY <= 1 {
		b.movementY *= -1
	}
}

type Game struct {
	bouncers    []Bouncer
	pressedKeys []ebiten.Key
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
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

	for _, key := range g.pressedKeys {
		switch key.String() {
		case "ArrowDown":
			if lineWidth > 0.25 {
				lineWidth -= 0.25
			}
		case "ArrowUp":
			if lineWidth < 50.0 {
				lineWidth += 0.25
			}
		}
	}

	for pos, bouncer := range g.bouncers {
		bouncer.update()
		g.bouncers[pos] = bouncer
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 1; i < len(g.bouncers); i++ {
		vector.StrokeLine(screen,
			float32(g.bouncers[i-1].positionX),
			float32(g.bouncers[i-1].positionY),
			float32(g.bouncers[i].positionX),
			float32(g.bouncers[i].positionY),
			lineWidth,
			g.bouncers[i].colour, true)
	}

	lastBouncer := g.bouncers[len(g.bouncers)-1]
	vector.StrokeLine(screen,
		float32(lastBouncer.positionX), float32(lastBouncer.positionY),
		float32(g.bouncers[0].positionX), float32(g.bouncers[0].positionY),
		lineWidth, lastBouncer.colour, true)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("This is it!")

	game := Game{}
	game.initBouncers()

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
