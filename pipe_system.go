package main

import (
    //_ "image/png"

    "github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
)

type PipeSystem struct {
    pipes [PIPE_NUM]*Pipe
    first_pipe_idx int
    last_pipe_idx int
}

func new_pipe_system() *PipeSystem {
    system := PipeSystem{}
    for a := 0; a < PIPE_NUM; a++ {
        system.pipes[a] = new_pipe()
    }
    system.first_pipe_idx = 0
    system.last_pipe_idx = PIPE_NUM - 1
    system.release_new_pipe()
    return &system
}

func (system *PipeSystem) draw(screen *ebiten.Image) {
    for _, pipe := range system.pipes {
        pipe.draw(screen)
    }
}

func (system *PipeSystem) move() {
    for _, pipe := range system.pipes {
        pipe.move()
    }
    if system.pipes[system.last_pipe_idx].longitude <= SCREEN_WIDTH - PIPE_WIDTH - DISTANCE {
        system.release_new_pipe()
    }
    if system.pipes[system.first_pipe_idx].longitude < -PIPE_WIDTH {
        system.pipes[system.first_pipe_idx].reset()
        system.first_pipe_idx += 1
        system.first_pipe_idx %= PIPE_NUM
    }
}

func (system *PipeSystem) release_new_pipe() {
    system.last_pipe_idx += 1
    system.last_pipe_idx %= PIPE_NUM
    system.pipes[system.last_pipe_idx].visible = true
}
