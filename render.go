package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Renderer struct {
	game *Game
}

func NewRenderer(game *Game) *Renderer {
	return &Renderer{game: game}
}

func (r *Renderer) Render() {
	ClearScreenWindows()
	h := r.game.gridHeight
	w := r.game.gridWidth

	grid := newGrid(h, w)
	renderGrid(grid, r.game.snake, r.game.apple)

}

func newGrid(h int, w int) [][]rune {
	grid := make([][]rune, w+2)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]rune, h+2)
	}
	return grid
}

func renderGrid(grid [][]rune, snake Snake, apple Apple) {
	drawBorders(grid)
	drawSnake(grid, snake)
	drawApple(grid, apple)
	printGrid(grid)
}

func printGrid(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", grid[i][j])
			}
		}
		fmt.Println()
	}
}

func drawSnake(grid [][]rune, snake Snake) {
	for i, pos := range snake.body {
		if i == 0 {
			grid[pos.y+1][pos.x+1] = '@'
		} else {
			grid[pos.y+1][pos.x+1] = 'o'
		}
	}
}

func drawApple(grid [][]rune, apple Apple) {
	grid[apple.location.y+1][apple.location.x+1] = 'M'
}

func drawBorders(grid [][]rune) {
	if len(grid) <= 0 {
		return //valutare se fixare
	}
	drawRow(grid[0], '_')           //bordo superiore
	drawRow(grid[len(grid)-1], '_') //bordo inferiore
	drawColumn(grid, 0, '|')
	drawColumn(grid, len(grid[0])-1, '|')
}

func drawRow(row []rune, ch rune) {
	for i := range row {
		row[i] = ch
	}
}

func drawColumn(grid [][]rune, x int, ch rune) {
	for i := 1; i < len(grid); i++ { // i parte da 1 perché _ è posizionato nella parte inferiore della riga.
		grid[i][x] = ch
	}
}

func ClearScreenWindows() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
