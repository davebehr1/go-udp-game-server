package gameObjects

import (
	"fmt"
	"image"
	_ "image/png"
	"net"
	"os"

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

type Client struct {
}

func (c Client) Run() {

	fmt.Println("running client")

	paddle := Paddle{NONE}

	paddle.LoadContent()

	fmt.Print(paddle)

	cfg := pixelgl.WindowConfig{
		Title:  "Game Networking",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := LoadPicture("gopher.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Wheat)

	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(0.15, 0.15))
	sprite.Draw(win, mat)
	fmt.Println("running")

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

	for !win.Closed() {

		if win.Pressed(pixelgl.KeyLeft) {
			buf := []byte("pressed left")
			_, err := Conn.Write(buf)
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
