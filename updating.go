package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

func (p *paddle) update(keyState []uint8, et float32) {
	if keyState[sdl.SCANCODE_UP] != 0 && p.y - p.h / 2 > 0 {
		p.y -= p.speed * et
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 && p.y + p.h / 2 < float32(winHeight) {
		p.y += p.speed * et
	}
}

func (p *paddle) aiUpdate(ball *ball, et float32) {
	if ball.y + p.h / 2 < float32(winHeight) && ball.y - p.h / 2 > 0 {
		p.y = ball.y
	}
}

func (b *ball) update(p1, p2 *paddle, et float32) {
	if detectCollision(b, p1) {
		if b.y < p1.y+p1.h/2 && b.y > p1.y-p1.h/2 && b.dx < 0{
			b.dx = -b.dx
		} else if b.x < p1.x + p1.w / 2 && b.x > p1.x - p1.w / 2 {
			b.dy = -b.dy
		} else {
			testX, testY := closestCorner(b, p1)
			vx, vy := testX - b.x, testY - b.y
			vd := float32(math.Sqrt(float64(vx * vx + vy * vy)))
			ux, uy := vx / vd, vy / vd
			b.x, b.y = testX - b.radius * ux, testY - b.radius * uy

			q := -(2 * (b.dx * (b.x - testX) + b.dy * (b.y - testY)) / (b.radius * b.radius))
			b.dx = b.dx + q * (b.x - testX)
			b.dy = b.dy + q * (b.y - testY)
		}
	}

	b.x += b.dx * et
	b.y += b.dy * et
	if b.y - b.radius < 1 || b.y + b.radius > float32(winHeight) - 1 {
		b.dy = -b.dy
	}

	if b.x - b.radius < 0 {
		p2.score++
		b.pos = getCenter()
		state = pause
	} else if b.x + b.radius > float32(winWidth) {
		p1.score++
		b.pos = getCenter()
		state = pause
	}

	if b.x + b.radius >= p2.x - p2.w / 2 && b.dx > 0{
		if b.y < p2.y+p2.h/2 && b.y > p2.y-p2.h/2 {
			b.dx = -b.dx
		}
	}
}