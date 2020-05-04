package main

import "github.com/veandco/go-sdl2/sdl"

func (p *paddle) draw(pixels []byte) {
	startX := int(p.x - p.w / 2)
	startY := int(p.y - p.h / 2)

	for y := 0; y < int(p.h); y++ {
		for x := 0; x < int(p.w); x++ {
			setPixel(startX + x, startY + y, p.color, pixels)
		}
	}

	numPos := pos{
		x: lerp(p.x, getCenter().x, 0.2),
		y: 50,
	}
	drawNumber(numPos, p.color, 10, p.score, pixels)
}

func (b *ball) draw(pixels []byte) {
	for y := -b.radius; y < b.radius; y++ {
		for x := -b.radius; x < b.radius; x++ {
			if x * x + y * y < b.radius * b.radius {
				setPixel(int(b.x+x), int(b.y+y), b.color, pixels)
			}
		}
	}
}

func drawNumber(pos pos, color sdl.Color, size, num int, pixels []byte) {
	startX := int(pos.x) - (size * 3) / 2
	startY := int(pos.y) - (size * 5) / 2

	for i, v := range nums[num] {
		if v == 1 {
			for y := startY; y < startY + size; y++ {
				for x := startX; x < startX + size; x++ {
					setPixel(x, y, color, pixels)
				}
			}
		}
		startX += size
		if (i + 1) % 3 == 0 {
			startY += size
			startX -= size * 3
		}
	}
}