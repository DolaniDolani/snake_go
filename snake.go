package main

type Snake struct {
	body []Pos
}

func (snake *Snake) MoveUpwards() {
	newHead := Pos{
		x: snake.body[0].x,
		y: snake.body[0].y - 1,
	}
	snake.body = append([]Pos{newHead}, snake.body...)
	snake.body = snake.body[:len(snake.body)-1]
}

func (snake *Snake) MoveDownwards() {
	newHead := Pos{
		x: snake.body[0].x,
		y: snake.body[0].y + 1,
	}
	snake.body = append([]Pos{newHead}, snake.body...)
	snake.body = snake.body[:len(snake.body)-1]
}

func (snake *Snake) MoveWestwards() {
	newHead := Pos{
		x: snake.body[0].x - 1,
		y: snake.body[0].y,
	}
	snake.body = append([]Pos{newHead}, snake.body...)
	snake.body = snake.body[:len(snake.body)-1]
}

func (snake *Snake) MoveEastwards() {
	newHead := Pos{
		x: snake.body[0].x + 1,
		y: snake.body[0].y,
	}
	snake.body = append([]Pos{newHead}, snake.body...)
	snake.body = snake.body[:len(snake.body)-1]
}

func (snake *Snake) Grow(newHead Pos) {
	snake.body = append([]Pos{newHead}, snake.body...)
}
