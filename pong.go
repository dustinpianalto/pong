package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

var (
	winWidth = 800
	winHeight = 600
	state = start
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"PONG",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		int32(winWidth),
		int32(winHeight),
		sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STREAMING,
		int32(winWidth),
		int32(winHeight))
	if err != nil {
		log.Fatal(err)
	}
	defer tex.Destroy()

	player1 := paddle{
		pos:   pos{50, 300},
		w:     20,
		h:     100,
		speed: 300,
		score: 0,
		color: sdl.Color{255, 0, 0, 255},
	}

	player2 := paddle{
		pos:   pos{750, 300},
		w:     20,
		h:     100,
		speed: 300,
		score: 0,
		color: sdl.Color{0, 255, 0, 255},
	}

	ball := ball{
		pos:    pos{400, 300},
		radius: 20,
		dx:     400,
		dy:     400,
		color:  sdl.Color{0, 0, 255, 255},
	}

	keyState := sdl.GetKeyboardState()
	pixels := make([]byte, winHeight*winWidth*4)

	resetBall(&ball)

	gameLoop(pixels, player1, player2, ball, renderer, tex, keyState)

}


