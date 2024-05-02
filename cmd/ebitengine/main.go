package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/therealfakemoot/automata"
)

type Game struct {
	width, height int
	pixels        []byte
	world         *automata.Grid
	rule          automata.Rule
}

func (g *Game) Update() error {
	g.world.Update(g.rule)
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			cell := g.world.Get(i, j)
			pixelIdx := i + (i * j)
			switch cell {
			case 0:
				g.pixels[pixelIdx*4] = 0
				g.pixels[pixelIdx*4+1] = 0
				g.pixels[pixelIdx*4+2] = 0
				g.pixels[pixelIdx*4+3] = 0
			case 1:
				g.pixels[pixelIdx*4] = 0xff
				g.pixels[pixelIdx*4+1] = 0xff
				g.pixels[pixelIdx*4+2] = 0xff
				g.pixels[pixelIdx*4+3] = 0xff
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func main() {
	gameWidth, gameHeight := 300, 300
	world := automata.NewGrid(gameWidth, gameHeight, automata.SeedCenter)
	g := &Game{
		width:  gameWidth,
		height: gameHeight,
		pixels: make([]byte,
			gameWidth*gameHeight*4),
		world: &world,
		rule:  automata.Rule30,
	}
	ebiten.SetWindowSize(gameWidth*2, gameHeight*2)
	ebiten.SetWindowTitle("Automata Demo")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("game error: %s\n", err)
	}
}
