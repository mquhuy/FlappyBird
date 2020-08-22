package main

import (
	_ "image/png"
	"github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
    "strconv"
)

type Score struct {
    score_images [10]*ebiten.Image
    points int
}

func initiate_scores() *Score {
    score := Score{}
    score.reset()
    for i := 0; i <= 9; i++ {
        img_name := strconv.Itoa(i)
        score.score_images[i], _, _ = ebitenutil.NewImageFromFile("images/" + img_name + ".png", ebiten.FilterDefault)
    }
    return &score
}

func(score *Score) draw_score(screen *ebiten.Image, center_x float64, center_y float64) {
    score_in_string := strconv.Itoa(score.points)
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(center_x - float64(len(score_in_string))*float64(NUM_IMG_WIDTH)/2, center_y - float64(NUM_IMG_HEIGHT)/2)
    for _, char := range score_in_string {
        value := string(char)
        value_num, _ := strconv.Atoi(value)
        screen.DrawImage(score.score_images[value_num], op)
        op.GeoM.Translate(NUM_IMG_WIDTH, 0)
    }
}

func(score *Score) reset() {
    score.points = 0
}
