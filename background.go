package main

import (
    _ "image/png"

    "github.com/hajimehoshi/ebiten" 
    "github.com/hajimehoshi/ebiten/ebitenutil"
)

type Background struct {
    background *ebiten.Image
    base *ebiten.Image
    latitude float64
}

func new_background() *Background {
    background := Background{}
    background.background, _, _ = ebitenutil.NewImageFromFile("images/background-night.png", ebiten.FilterDefault)
    background.base, _, _ = ebitenutil.NewImageFromFile("images/base.png", ebiten.FilterDefault)
    background.latitude = 0
    return &background
}

func (background *Background) move() {
    background.latitude -= VELOCITY
    if background.latitude <= -BG_WIDTH {
        background.latitude += BG_WIDTH
    }
}

func (background *Background) draw(screen *ebiten.Image) {
	op_bg := &ebiten.DrawImageOptions{}
	op_base := &ebiten.DrawImageOptions{}
    op_base.GeoM.Translate(background.latitude, BG_HEIGHT)
    screen.DrawImage(background.background, op_bg)
    screen.DrawImage(background.base, op_base)
    for i := 1; i < BG_NUM; i++ {
        op_bg.GeoM.Translate(BG_WIDTH, 0)
        screen.DrawImage(background.background, op_bg)
        op_base.GeoM.Translate(BG_WIDTH, 0)
        screen.DrawImage(background.base, op_base)
    }
    op_base.GeoM.Translate(BG_WIDTH, 0)
    screen.DrawImage(background.base, op_base)
}
