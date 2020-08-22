package main

import (
    _ "image/png"

    "github.com/hajimehoshi/ebiten"
)

func draw_in_center(img, screen *ebiten.Image) {
    img_width, img_height := img.Size()
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(float64(SCREEN_WIDTH - img_width)/2, float64(SCREEN_HEIGHT - img_height)/2)
    screen.DrawImage(img, op)
}
