package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	screenW = 300
	screenH = 300
	originX = 0
	originY = 0
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Maze",
		Bounds: pixel.R(0, 0, screenW, screenH),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	exitOnError(err)

	data := []*imdraw.IMDraw{}

	var x float64 = 0
	var y float64 = 150
	var size float64 = 20
	for i := 0; i < int(screenW-size); i++ {
		data = append(data, drawRectangle(x, y, x+size, y+size, colornames.Blue))
		x += 1
	}

	j := 0
	forward := true
	delay := 5
	frame := 0
	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyQ) {
			exit()
		}
		win.Clear(colornames.Black)
		data[j].Draw(win)
		if frame < delay {
			frame++
		} else {
			frame = 0
			if forward {
				if j < len(data)-1 {
					j++
				} else {
					forward = false
				}
			} else {
				if j > 1 {
					j--
				} else {
					forward = true
				}
			}
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func exit() {
	os.Exit(0)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func drawRectangle(x, y, x2, y2 float64, c color.Color) *imdraw.IMDraw {
	s := imdraw.New(nil)
	s.Color = c
	s.Push(pixel.V(x, y))
	s.Push(pixel.V(x2, y2))
	s.Rectangle(0)
	return s
}
