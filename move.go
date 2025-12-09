package main

type MoveResult int

const (
	MoveInvalid MoveResult = iota
	MoveOk
	MoveEat
	MoveDeath
)
