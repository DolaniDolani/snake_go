package main

import "github.com/eiannone/keyboard"

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	startGame()
}
