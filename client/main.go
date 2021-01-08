package main

import (
	"fmt"

	"github.com/davebehr1/goUDP/gameObjects"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	c := gameObjects.Client{}
	fmt.Println("Running Client")
	pixelgl.Run(c.Run)
}
