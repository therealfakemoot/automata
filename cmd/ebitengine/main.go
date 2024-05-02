package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	width, height int
	pixels        []byte
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outwideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func main() {
	gameWidth, gameHeight := 300, 300
	g := &Game{width: gameWidth, height: gameHeight, pixels: make([]byte, gameWidth*gameHeight*4)}
	ebiten.SetWindowSize(gameWidth*2, gameHeight*2)
	ebiten.SetWindowTitle("Automata Demo")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("game error: %s\n", err)
	}
}
