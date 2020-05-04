package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"time"
)

func closestCorner(b *ball, p *paddle) (float32, float32) {
	testX := b.x
	testY := b.y
	if b.x < p.x - p.w / 2 {
		testX = p.x - p.w / 2
	} else if b.x > p.x + p.w / 2 {
		testX = p.x + p.w / 2
	}
	if b.y < p.y - p.h / 2 {
		testY = p.y - p.h / 2
	} else if b.y > p.y + p.h / 2 {
		testY = p.y + p.h / 2
	}
	return testX, testY
}

func detectCollision(b *ball, p *paddle) bool {
	testX, testY := closestCorner(b, p)

	sideX := b.x - testX
	sideY := b.y - testY
	return b.radius * b.radius >= sideX * sideX + sideY * sideY
}

func setPixel(x, y int, c sdl.Color, pixels []byte) {
	index := (y*winWidth + x) * 4
	if index < len(pixels) - 4 && index >= 0 {
		pixels[index] = c.R
		pixels[index+1] = c.G
		pixels[index+2] = c.B
		//pixels[index+3] = c.A
	}
}

func clearScreen(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func getCenter() pos {
	return pos{float32(winWidth) / 2, float32(winHeight) / 2}
}

func lerp(a, b, pct float32) float32 {
	return a + pct * (b - a)
}

func resetBall(b *ball) {
	choices := []float32{-400, -200, 200, 400}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	b.dy = choices[r.Intn(len(choices))]
	b.dx = choices[r.Intn(len(choices))]
}