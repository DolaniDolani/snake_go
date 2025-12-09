package main

import "math/rand"

type Pos struct {
	x int
	y int
}

func RandomPosition(maxWidth int, maxHeight int) Pos {
	x := rand.Intn(maxWidth)
	y := rand.Intn(maxHeight)
	return Pos{x: x, y: y}
}
