package main

import (
	"fmt"
	"time"
)

type Game struct {
	gridHeight int
	gridWidth  int
	snake      Snake
	apple      Apple
	lost       bool
	lastMove   rune
}

func (game *Game) MoveSnake(input rune) {
	switch input {
	case 'w':
		if game.snake.body[0].y-1 == game.snake.body[1].y {
			break
		}
		game.snake.MoveUpwards()
	case 'a':
		if game.snake.body[0].x-1 == game.snake.body[1].x {
			break
		}
		game.snake.MoveWestwards()
	case 's':
		if game.snake.body[0].y+1 == game.snake.body[1].y {
			break
		}
		game.snake.MoveDownwards()
	case 'd':
		if game.snake.body[0].x+1 == game.snake.body[1].x {
			break
		}
		game.snake.MoveEastwards()
	}
}

func (game *Game) NextPosition(input rune) Pos {
	switch input {
	case 'w':
		return Pos{x: game.snake.body[0].x, y: game.snake.body[0].y - 1}
	case 'a':
		return Pos{x: game.snake.body[0].x - 1, y: game.snake.body[0].y}
	case 's':
		return Pos{x: game.snake.body[0].x, y: game.snake.body[0].y + 1}
	case 'd':
		return Pos{x: game.snake.body[0].x + 1, y: game.snake.body[0].y}
	default:
		return Pos{0, 0}
	}
}

func (game *Game) EvaluateMove(input rune) MoveResult {
	//se testa su mela -> mela scompare, cresciamo, appare nuova mela.
	nextHead := game.NextPosition(input)
	if game.wouldEatApple(nextHead) {
		return MoveEat
	}
	//se testa su corpo -> perdiamo
	//se ci si prova a muovere indietro -> non ci si muove
	if game.wouldMoveBack(nextHead) {
		return MoveInvalid
	}
	//se testa su muro -> perdiamo
	if game.wouldClashWithBody(nextHead) || game.wouldClashWithBorder(nextHead) {
		return MoveDeath
	}
	//se stato "neutrale" -> ci muoviamo
	return MoveOk
}

func (game *Game) ActMove(input rune, moveResult MoveResult) {
	switch moveResult {
	case MoveDeath:
		fmt.Println("Hai perso!")
		game.lost = true
	case MoveEat:
		game.EatApple()
		game.GenerateNewApple()
		game.lastMove = input
	case MoveOk:
		game.MoveSnake(input)
		game.lastMove = input
	case MoveInvalid:

	default:
		fmt.Println("INVALID MOVE STATE")
	}
}

func (game *Game) wouldClashWithBorder(nextPos Pos) bool {
	if nextPos.x < 0 || nextPos.y < 0 ||
		nextPos.x >= game.gridWidth || nextPos.x >= game.gridHeight {
		return true

	} else {
		return false
	}
}

func (game *Game) wouldClashWithBody(nextPos Pos) bool {
	for _, bodyPos := range game.snake.body {
		if bodyPos.x == nextPos.x && bodyPos.y == nextPos.y {
			return true
		}
	}
	return false
}

func (game *Game) wouldMoveBack(nextPos Pos) bool {
	if nextPos.x == game.snake.body[1].x && nextPos.y == game.snake.body[1].y {
		return true
	} else {
		return false
	}
}

func (game *Game) wouldEatApple(nextPos Pos) bool {
	if nextPos.x == game.apple.location.x &&
		nextPos.y == game.apple.location.y {
		return true
	} else {
		return false
	}
}

func (game *Game) EatApple() {
	game.snake.Grow(game.apple.location)
}

func (game *Game) GenerateNewApple() {
	validPosition := false
	apple := Apple{location: RandomPosition(game.gridWidth, game.gridHeight)}
	for !validPosition {
		validPosition = true
		for _, pos := range game.snake.body {
			if pos.x == apple.location.x && pos.y == apple.location.y {
				apple.location = RandomPosition(game.gridWidth, game.gridHeight)
				validPosition = false
				break
			}
		}
	}
	game.apple = apple
}

func newGame() Game {
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
		lost:     false,
		lastMove: 'd',
	}
	return game
}

func startGame() {
	game := newGame()
	timedLoop(&game)
}

func loop(game *Game) {
	renderer := NewRenderer(game)
	renderer.Render()

	for !game.lost {
		//input := ReadInput() //terminale normale
		input := ReadRawInputWindows()
		evaluatedMove := game.EvaluateMove(input)
		game.ActMove(input, evaluatedMove)
		renderer.Render()
	}
	fmt.Println("Grazie per aver giocato!")
}

func timedLoop(game *Game) {
	renderer := NewRenderer(game)
	renderer.Render()

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	inputChan := make(chan rune, 10)
	go func() {
		for {
			input := ReadRawInputWindows()
			inputChan <- input
		}
	}()
	var pendingInput rune = game.lastMove
	for !game.lost {

		select {
		case input := <-inputChan:
			pendingInput = input
		case <-ticker.C:
			move := pendingInput
			evaluatedMove := game.EvaluateMove(move)
			game.ActMove(move, evaluatedMove)
			renderer.Render()
			game.lastMove = move
		}

	}
	fmt.Println("Grazie per aver giocato!")
}
