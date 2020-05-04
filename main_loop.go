package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

func gameLoop(pixels []byte, player1, player2 paddle, ball ball, renderer *sdl.Renderer, tex *sdl.Texture, keyState []uint8) {
	var frameStart time.Time
	var elapsedTime float32

	for {
		frameStart = time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		if err := renderer.Clear(); err != nil {
			log.Fatal(err)
		}

		if state == play {
			if player1.score == 5 || player2.score == 5 {
				state = start
			}
			player1.update(keyState, elapsedTime)
			player2.aiUpdate(&ball, elapsedTime)
			ball.update(&player1, &player2, elapsedTime)
		} else if state == start {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				player1.score = 0
				player2.score = 0
				resetBall(&ball)
				state = play
			}
		} else if state == pause {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				resetBall(&ball)
				state = play
			}
		}

		clearScreen(pixels)
		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)

		if err := tex.Update(nil, pixels, winWidth*4); err != nil {
			log.Fatal(err)
		}
		if err := renderer.Copy(tex, nil, nil); err != nil {
			log.Fatal(err)
		}
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())
		if elapsedTime < 0.005 {
			sdl.Delay(5 - uint32(elapsedTime / 1000))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}
}