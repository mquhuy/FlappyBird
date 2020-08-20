package main

import (
	_ "image/png"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
    "math"
)

type Bird struct {
    images [frameNum]*ebiten.Image
    longitude float64
    latitude float64
    alive bool
    active int
    count int
    idx_increment int
    velocity float64
}

func new_bird() *Bird {
    bird := Bird{}
    bird_upflap, _, _ := ebitenutil.NewImageFromFile("images/yellowbird-upflap.png", ebiten.FilterDefault)
    bird_midflap, _, _ := ebitenutil.NewImageFromFile("images/yellowbird-midflap.png", ebiten.FilterDefault)
    bird_downflap, _, _ := ebitenutil.NewImageFromFile("images/yellowbird-downflap.png", ebiten.FilterDefault)
    bird.images = [frameNum]*ebiten.Image{bird_upflap, bird_midflap, bird_downflap}
    bird.idx_increment = 1
    bird.reset()
    return &bird
}

func(bird *Bird) die() {
    bird.alive = false
}

func(bird *Bird) drop() {
    bird.latitude += bird.velocity
    if bird.latitude >= BG_HEIGHT - BIRD_HEIGHT {
        bird.latitude = BG_HEIGHT - BIRD_HEIGHT
        bird.velocity = 0
    } else {
        bird.velocity += GRAVITY
    }
}

func(bird *Bird) draw(screen *ebiten.Image) {
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(bird.longitude, bird.latitude)
    if !bird.alive {
        op.GeoM.Rotate(math.Pi/2)
        op.GeoM.Translate(2, 0)
    }
    screen.DrawImage(bird.images[bird.active], op)
}

func(bird *Bird) flap() {
    bird.count += bird.idx_increment
    bird.active = bird.count/5
    if (bird.count >= (frameNum-1)*5 || bird.count <= 0) {
        bird.idx_increment = -bird.idx_increment
    }
}

func(bird *Bird) jump() {
    bird.velocity -= BIRD_JUMP_ACC
}

func(bird *Bird) touch_pipe(pipe *Pipe) bool {
    if pipe.longitude < -PIPE_WIDTH || pipe.longitude > BIRD_WIDTH {
        return false
    }
    if bird.latitude <= pipe.top_height || bird.latitude + BIRD_HEIGHT >= pipe.top_height + GAP {
        return true
    }
    return false
}

func(bird *Bird) reset() {
    bird.alive = true
    bird.latitude = float64(BG_HEIGHT)/2 - float64(BIRD_HEIGHT)/2
    bird.longitude = 10
    bird.velocity = 0
}
