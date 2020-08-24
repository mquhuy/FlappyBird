package main

import (
    _ "image/png"

    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
    "github.com/hajimehoshi/ebiten/inpututil"
)

type Game struct {
    background *Background
    pipes [PIPE_NUM]*Pipe
    bird *Bird
    score *Score
    first_pipe_idx int
    last_pipe_idx int
    mode string
    game_over_sign *ebiten.Image
}

func new_game() *Game {
    game := Game{}
    game.mode = "waiting"
    game.score = initiate_scores()
    game.background = new_background()
    for a := 0; a < PIPE_NUM; a++ {
        game.pipes[a] = new_pipe()
    }
    game.bird = new_bird()
    game.game_over_sign, _, _ = ebitenutil.NewImageFromFile("images/gameover.png", ebiten.FilterDefault)
    game.reset()
    return &game
}

func(game *Game) start_game() {
    game.mode = "on"
    game.bird.active = true
    game.release_new_pipe()
}

func (game *Game) Update(screen *ebiten.Image) error {
    game.switch_mode()
    game.bird.update()
    switch game.mode {
    case "waiting":
        game.background.move()
    case "on":
        game.background.move()
        for _, pipe := range game.pipes {
            pipe.move()
        }
        if game.pipes[game.last_pipe_idx].longitude <= SCREEN_WIDTH - PIPE_WIDTH - DISTANCE {
            game.release_new_pipe()
        }
        if game.first_pipe().longitude < -PIPE_WIDTH {
            game.reset_first_pipe()
            game.score.points += 1
        }
        if game.bird.touch_pipe(game.first_pipe()) {
            game.mode = "over"
            game.bird.die()
        }
    case "over":
        game.bird.drop()
    }
	return nil
}

func (game *Game) switch_mode() {
    if (inpututil.IsKeyJustPressed(ebiten.KeySpace)) {
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
    game.background.draw(screen)
    for _, pipe := range game.pipes {
        pipe.draw(screen)
    }
    game.bird.draw(screen)
    switch game.mode {
    case "on":
        game.score.draw_score(screen, float64(SCREEN_WIDTH)/2, float64(SCREEN_HEIGHT)/3)
    case "over":
        draw_in_center(game.game_over_sign, screen)
    }
}

func (game *Game) reset() {
    for a := 0; a < PIPE_NUM; a++ {
        game.pipes[a].reset()
    }
    game.first_pipe_idx = 0
    game.last_pipe_idx = PIPE_NUM - 1
    game.mode = "waiting"
    game.bird.reset()
    game.score.reset()
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
