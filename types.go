package main

import "github.com/veandco/go-sdl2/sdl"

type pos struct {
	x, y float32
}

type ball struct {
	pos
	radius 		float32
	dx 			float32
	dy 			float32
	color 		sdl.Color
}

type paddle struct {
	pos
	w 			float32
	h 			float32
	speed		float32
	score 		int
	color 		sdl.Color
}

type gameState int
const (
	start gameState = iota
	play
	pause
	end
)