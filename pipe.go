package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
    "image"
    "math"
    "math/rand"
)

type Pipe struct {
    longitude float64
    top_height float64
    bot_pipe *ebiten.Image
    visible bool
}

func (pipe *Pipe) reset() {
    left_limit := float64(SCREEN_HEIGHT - PIPE_HEIGHT - GAP)
    pipe.top_height = rand.Float64()*(PIPE_HEIGHT - left_limit) + left_limit
    pipe.longitude = SCREEN_WIDTH + PIPE_WIDTH
    pipe.visible = false
}

func new_pipe() *Pipe {
    pipe := Pipe{}
    pipe.bot_pipe, _, _ = ebitenutil.NewImageFromFile("images/pipe-red.png", ebiten.FilterDefault)
    pipe.reset()
    return &pipe
}

func (pipe *Pipe) draw(screen *ebiten.Image) {
    // Top pipe
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Rotate(math.Pi)
    op.GeoM.Translate(pipe.longitude + PIPE_WIDTH, pipe.top_height)
    screen.DrawImage(pipe.bot_pipe, op)

    // Bot pipe
    op = &ebiten.DrawImageOptions{}
    op.GeoM.Translate(pipe.longitude, pipe.top_height + GAP)
    bot_pipe_img := pipe.bot_pipe.SubImage(image.Rect(0, 0, PIPE_WIDTH, int(BG_HEIGHT - pipe.top_height - GAP + 3))).(*ebiten.Image)
    screen.DrawImage(bot_pipe_img, op)
}

func (pipe *Pipe) move(){
    if pipe.visible {
        pipe.longitude -= VELOCITY
    }
}
