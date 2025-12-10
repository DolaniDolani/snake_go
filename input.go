package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func ReadInput() rune {
	var input rune
	fmt.Scanf("%c ", &input)
	switch input {
	case 'w', 'a', 's', 'd':
		return input
	default:
		return 0
	}
}

func ReadRawInputWindows() rune {
	//defer keyboard.Close()
	char, _, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	return char
}
