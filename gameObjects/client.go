package gameObjects

import (
	"fmt"
	"image"
	_ "image/png"
	"net"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type PaddleSide int

const (
	NONE PaddleSide = iota
	LEFT
	RIGHT
)

const serverAddress = "localhost:1337"

var cfg pixelgl.WindowConfig

type Client struct {
}

func (c Client) Run() {

	fmt.Println("running client")

	cfg = pixelgl.WindowConfig{
		Title:  "Game Networking",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Black)

	leftPaddle := newPaddle(LEFT, cfg)

	rightPaddle := newPaddle(RIGHT, cfg)

	ball := newBall(cfg)

	ServerAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		panic(err)
		fmt.Println("panic")
	}
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		panic(err)
	}

	defer Conn.Close()

	last := time.Now()
	for !win.Closed() {

		dt := time.Since(last).Seconds()

		ctrl := pixel.ZV
		if win.Pressed(pixelgl.KeyUp) {
			ctrl.Y = 1
		}
		if win.Pressed(pixelgl.KeyDown) {
			ctrl.Y = -1
		}

		win.Clear(colornames.Black)

		leftPaddle.update(dt, ctrl)
		rightPaddle.update(dt, ctrl)

		leftPaddle.Draw(win)
		rightPaddle.Draw(win)

		ball.update(rightPaddle, leftPaddle, cfg)
		ball.Draw(win)

		if win.Pressed(pixelgl.KeyLeft) {
			packet := Packet{
				"pressed left",
				time.Now(),
				PaddlePosition,
			}
			_, err = Conn.Write(packet.getBytes())
			if err != nil {
				fmt.Println(err)
			}
		}
		win.Update()

	}
}

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
