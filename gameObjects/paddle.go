package gameObjects

import (
	"fmt"

	"github.com/faiface/pixel"
)

// var Paddle = "paddle"

var sprite *pixel.Sprite

type Paddle struct {
	side PaddleSide
}

func (p Paddle) LoadContent() {
	// pic, err := loadPicture("gopher.png")
	// _sprite := pixel.NewSprite(pic, pic.Bounds())
	fmt.Println("loading content")
}
