package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"math/rand"
)

var (
	circleShader *ebiten.Shader
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

// Improved draw and fill circle from https://golangprojectstructure.com/creating-cool-games-with-ebiten-in-go/
// This time, using shaders! Because really, this is where we want to be.
func (g *Game) drawCircle(screen *ebiten.Image, x, y, radius int, clr color.RGBA) {
	var path vector.Path

	path.MoveTo(float32(x), float32(y))
	path.Arc(float32(x), float32(y), float32(radius), 0, math.Pi*2, vector.Clockwise)

	vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)

	redScaled := float32(clr.R) / 255
	greenScaled := float32(clr.G) / 255
	blueScaled := float32(clr.B) / 255
	alphaScaled := float32(clr.A) / 255

	for i := range vertices {
		v := &vertices[i]

		v.ColorR = redScaled
		v.ColorG = greenScaled
		v.ColorB = blueScaled
		v.ColorA = alphaScaled
	}

	screen.DrawTrianglesShader(vertices, indices, circleShader, &ebiten.DrawTrianglesShaderOptions{
		FillRule: ebiten.EvenOdd,
	})
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, bouncer := range g.bouncers {
		x := int(math.Round(bouncer.positionX))
		y := int(math.Round(bouncer.positionY))
		g.drawCircle(screen, x, y, bouncerSize, bouncer.colour)
	}
}

func init() {
	var err error

	circleShader, err = ebiten.NewShader([]byte(`
		package main

		func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
			return color
		}
	`))
	if err != nil {
		panic(err)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Rounders all around!")

	game := Game{}
	game.initBouncers()

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
