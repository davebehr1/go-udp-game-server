package gameObjects

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Ball struct {
	pixel.Circle
	velocity pixel.Vec
	color    color.Color
}

func newBall(cfg pixelgl.WindowConfig) *Ball {
	ball := &Ball{
		pixel.C(cfg.Bounds.Center(), 10.0),
		pixel.V(8, 0),
		colornames.White,
	}

	return ball
}

func (b *Ball) Draw(win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Push(b.Center)
	imd.Circle(b.Radius, 0)
	imd.Draw(win)
}

func (b *Ball) update(rPaddle *Paddle, lPaddle *Paddle, cfg pixelgl.WindowConfig) {

	if b.Center.X+b.Radius >= rPaddle.pos.X-rPaddle.w/2 &&
		(b.Center.Y < rPaddle.pos.Y+rPaddle.h/2 && b.Center.Y > rPaddle.pos.Y-rPaddle.h/2) {
		b.Center.X -= rPaddle.w/2 + b.Radius
		b.velocity.X = -b.velocity.X
		b.velocity.Y = rPaddle.areaHit(b)
	}
	// Check collision with left paddle
	if b.Center.X-b.Radius <= lPaddle.pos.X+lPaddle.w/2 &&
		(b.Center.Y < lPaddle.pos.Y+lPaddle.h/2 && b.Center.Y > lPaddle.pos.Y-lPaddle.h/2) {
		b.Center.X += lPaddle.w/2 + b.Radius
		b.velocity.X = -b.velocity.X
		b.velocity.Y = lPaddle.areaHit(b)
	}

	// left wall
	if b.Center.X <= 0 {
		b.Center = cfg.Bounds.Center()
	}
	//right wall
	if b.Center.X >= cfg.Bounds.Max.X {
		b.Center = cfg.Bounds.Center()
	}

	//top and bottom
	if b.Center.Y-b.Radius < 0 || b.Center.Y+b.Radius > cfg.Bounds.Max.Y {
		b.velocity.Y = -b.velocity.Y
	}
	b.Center = b.Center.Add(b.velocity)

}
