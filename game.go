package main

import "fmt"

type Game struct {
	gridHeight int
	gridWidth  int
	snake      Snake
	apple      Apple
}

func createGame() Game {
	game := Game{
		gridWidth:  15,
		gridHeight: 15,
		snake: Snake{
			[]Pos{
				{x: 7, y: 7},
				{x: 6, y: 7},
			},
		},
		apple: Apple{
			Pos{x: 9, y: 7},
		},
	}
	return game
}

func startGame() {
	game := createGame()
	loop(&game)
}

func loop(game *Game) {
	renderer := NewRenderer(game)
	renderer.Render()
	readInput()
}

func readInput() rune {
	var input rune
	fmt.Scanf("%c", &input)
	switch input {
	case 'w', 'a', 's', 'd':
		fmt.Println("input valido: movimento")
		return input
	default:
		fmt.Println("input non valido")
		return 0
	}
}
