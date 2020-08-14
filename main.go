package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
    // Height and Width of the background image
    BG_WIDTH = 288
	BG_HEIGHT = 512

    // Number of background images stacking together
    BG_NUM = 4

    // Height of the base.png image
    BASE_HEIGHT = 112

    // Height and Width of the screen
    SCREEN_HEIGHT = BG_HEIGHT+BASE_HEIGHT
	SCREEN_WIDTH  = BG_WIDTH*BG_NUM

    // Number of different bird images
	frameNum = 3

    // GRAVITY of the earth
    GRAVITY = 10

    // Height and width of the bird image
    BIRD_HEIGHT = 24
    BIRD_WIDTH = 34

    // Height and width of the pipe image
    PIPE_HEIGHT = 320
    PIPE_WIDTH = 52

    // The gap between the top and bot pipes
    GAP = 150

    // Velocity per frame of the bird
    VELOCITY = 2

    // Distance between two pipes
    DISTANCE = 150

    // Number of pipes needed
    PIPE_NUM = SCREEN_WIDTH / DISTANCE + 1
)

var (
	background, base *ebiten.Image
)

func main() {
    img, _, err := ebitenutil.NewImageFromFile("images/background-night.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
    background = img
    base, _, err = ebitenutil.NewImageFromFile("images/base.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
    game := new_game()

	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Flappy Bird")
    //ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
