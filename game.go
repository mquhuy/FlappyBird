package main

import (
    _ "image/png"

    "github.com/hajimehoshi/ebiten" 
    "github.com/hajimehoshi/ebiten/ebitenutil"
    "github.com/hajimehoshi/ebiten/inpututil"
)

const (
    KeySpace = iota
)

type Game struct {
    pipes [PIPE_NUM]*Pipe
    bird *Bird
    point int
    first_pipe_idx int
    last_pipe_idx int
    background *ebiten.Image
    base *ebiten.Image
    mode string
}

func new_game() *Game {
    game := Game{}
    game.mode = "waiting"
    game.background, _, _ = ebitenutil.NewImageFromFile("images/background-night.png", ebiten.FilterDefault)
    game.base, _, _ = ebitenutil.NewImageFromFile("images/base.png", ebiten.FilterDefault)
    for a := 0; a < PIPE_NUM; a++ {
        game.pipes[a] = new_pipe()
    }
    game.bird = new_bird()
    game.reset()
    return &game
}

func(game *Game) start_game() {
    game.mode = "on"
    game.release_new_pipe()
}

func (game *Game) Update(screen *ebiten.Image) error {
    game.switch_mode()
    switch game.mode {
    case "waiting":
        game.bird.flap()
    case "on":
        game.bird.flap()
        game.bird.drop()
        for _, pipe := range game.pipes {
            pipe.move()
        }
        if game.pipes[game.last_pipe_idx].longitude <= SCREEN_WIDTH - PIPE_WIDTH - DISTANCE {
            game.release_new_pipe()
        }
        if game.first_pipe().longitude < -PIPE_WIDTH {
            game.reset_first_pipe()
            game.point += 1
            print(game.point)
        }
        if game.bird.touch_pipe(game.first_pipe()) {
            game.mode = "over"
        }
    case "over":
        game.bird.die()
    }
	return nil
}

func (game *Game) switch_mode() {
    if (inpututil.IsKeyJustPressed(KeySpace)) {
        switch game.mode {
        case "waiting":
            game.start_game()
        case "on":
            game.bird.jump()
        case "over":
            game.reset()
        }
    }
}

func (game *Game) Draw(screen *ebiten.Image) {

	op_bg := &ebiten.DrawImageOptions{}
	op_base := &ebiten.DrawImageOptions{}
    op_base.GeoM.Translate(0, BG_HEIGHT)
    screen.DrawImage(game.background, op_bg)
    screen.DrawImage(game.base, op_base)
    for i := 1; i < BG_NUM; i++ {
        op_bg.GeoM.Translate(BG_WIDTH, 0)
        op_base.GeoM.Translate(BG_WIDTH, 0)
        screen.DrawImage(game.background, op_bg)
        screen.DrawImage(game.base, op_base)
    }
    for _, pipe := range game.pipes {
        pipe.draw(screen)
    }
    game.bird.draw(screen)
}

func (game *Game) reset() {
    for a := 0; a < PIPE_NUM; a++ {
        game.pipes[a].reset()
    }
    game.first_pipe_idx = 0
    game.last_pipe_idx = PIPE_NUM - 1
    game.mode = "waiting"
    game.point = 0
    game.bird.reset()
}

func (game *Game) check_input() {
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}


func (game *Game) release_new_pipe() {
    game.last_pipe_idx += 1
    game.last_pipe_idx %= PIPE_NUM
    game.last_pipe().visible = true
}

func (game *Game) reset_first_pipe() {
    game.first_pipe().reset()
    game.first_pipe_idx += 1
    game.first_pipe_idx %= PIPE_NUM
}

func (game *Game) first_pipe() *Pipe {
    return game.pipes[game.first_pipe_idx]
}

func (game *Game) last_pipe() *Pipe {
    return game.pipes[game.last_pipe_idx]
}
