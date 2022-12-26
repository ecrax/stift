package main

import (
	"fmt"
	"os"
)

const WIDTH = 600
const HEIGHT = 400

type Canvas struct {
	pixels        []uint
	width, height int
}

func (canvas *Canvas) fill(color uint) {
	for y := 0; y < canvas.height; y++ {
		for x := 0; x < canvas.width; x++ {
			canvas.pixels[y*canvas.width+x] = color
		}
	}
}

func initCanvas(width int, height int) *Canvas {
	return &Canvas{make([]uint, width*height), width, height}
}

func main() {
	canvas := initCanvas(WIDTH, HEIGHT)
	fmt.Println(canvas.pixels)
	canvas.fill(0xFFFFFFFF)
	fmt.Println(canvas.pixels)

	f, err := os.Create("./fill.ppm")

}
