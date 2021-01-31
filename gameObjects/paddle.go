package gameObjects

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type gameObject interface {
	update()
	Draw()
}
type Paddle struct {
	side  PaddleSide
	pos   pixel.Vec
	Speed float64
	w, h  float64
	color color.Color
}

func newPaddle(side PaddleSide, cfg pixelgl.WindowConfig) *Paddle {

	var pos pixel.Vec
	if side == LEFT {
		pos = pixel.V(cfg.Bounds.Max.X/2, cfg.Bounds.Max.Y/2).Sub(pixel.V(cfg.Bounds.Max.X/2-20, 0))
	} else if side == RIGHT {
		pos = pixel.V(cfg.Bounds.Max.X/2, cfg.Bounds.Max.Y/2).Add(pixel.V(cfg.Bounds.Max.X/2-20, 0))
	}

	paddle := &Paddle{
		LEFT,
		pos,
		5.0,
		10,
		70,
		colornames.White,
	}
	return paddle
}

func (p *Paddle) rect() pixel.Rect {
	return pixel.R(p.pos.X-p.w/2, p.pos.Y-p.h/2, p.pos.X+p.w/2, p.pos.Y+p.h/2)
}

func (pP *Paddle) update(dt float64, ctrl pixel.Vec) {
	switch {
	case ctrl.Y > 0:
		pP.pos.Y += pP.Speed
	case ctrl.Y < 0:
		pP.pos.Y -= pP.Speed
	}
}
func (p *Paddle) areaHit(ball *Ball) float64 {
	var a, f pixel.Vec
	if p.side == RIGHT {
		f, a = p.rect().Vertices()[0], p.rect().Vertices()[1]
	} else if p.side == LEFT {
		a, f = p.rect().Vertices()[2], p.rect().Vertices()[3]
	}
	b := pixel.Lerp(a, f, .2)
	c := pixel.Lerp(a, f, .4)
	d := pixel.Lerp(a, f, .6)
	e := pixel.Lerp(a, f, .8)

	switch y := ball.Center.Y; {
	case y < a.Y && y > b.Y:
		return 10
	case y < b.Y && y > c.Y:
		return 5
	case y < c.Y && y > d.Y:
		return ball.velocity.Y
	case y < d.Y && y > e.Y:
		return -5
	case y < e.Y && y > f.Y:
		return -10
	default:
		return ball.velocity.Y
	}
}

func (p *Paddle) Draw(win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = p.color
	imd.Push(p.rect().Min)
	imd.Push(p.rect().Max)
	imd.Rectangle(0)
	imd.Draw(win)
}
