package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

const (
	screenWidth  = 320
	screenHeight = 240
)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple Snake")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

}
