package main

import (
	_ "image/png"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
    "github.com/lthh91/WAVPlayer"
    "math"
)

type Bird struct {
    images [frameNum]*ebiten.Image
    longitude float64
    latitude float64
    alive bool
    active bool
    active_img_idx int
    count int
    idx_increment int
    velocity float64
    swoosh *wavPlayer.WavPlayer
}

func new_bird() *Bird {
    bird := Bird{}
    bird_upflap, _, _ := ebitenutil.NewImageFromFile("images/yellowbird-upflap.png", ebiten.FilterDefault)
    bird_midflap, _, _ := ebitenutil.NewImageFromFile("images/yellowbird-midflap.png", ebiten.FilterDefault)
    bird_downflap, _, _ := ebitenutil.NewImageFromFile("images/yellowbird-downflap.png", ebiten.FilterDefault)
    bird.images = [frameNum]*ebiten.Image{bird_upflap, bird_midflap, bird_downflap}
    bird.idx_increment = 1
    bird.swoosh, _ = wavPlayer.NewWAV("audio/swoosh.wav")
    bird.reset()
    return &bird
}

func(bird *Bird) die() {
    bird.alive = false
    bird.active = false
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
    if !bird.alive {
        op.GeoM.Rotate(math.Pi/3)
        op.GeoM.Translate(BIRD_HEIGHT, 0)
    }
    op.GeoM.Translate(bird.longitude, bird.latitude)
    screen.DrawImage(bird.images[bird.active_img_idx], op)
}

func(bird *Bird) flap() {
    bird.count += bird.idx_increment
    bird.active_img_idx = bird.count/4
    if (bird.count >= (frameNum-1)*4 || bird.count <= 0) {
        bird.idx_increment = -bird.idx_increment
    }
}

func(bird *Bird) jump() {
    bird.swoosh.Play()
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

func(bird *Bird) update() {
    if bird.alive {
        if bird.active {
            bird.flap()
            bird.drop()
        } else {
            bird.flap()
        }
    } else {
        bird.drop()
    }
}

func(bird *Bird) reset() {
    bird.alive = true
    bird.active = false
    bird.latitude = float64(BG_HEIGHT)/2 - float64(BIRD_HEIGHT)/2
    bird.longitude = 20
    bird.velocity = 0
}
