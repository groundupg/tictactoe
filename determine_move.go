package main

import "math/rand"

type Assessed struct {
	score float64
	move  [2]Player
}

func DetermineMove(b Board, p Player) [2]Player {
	if p == 1 {
		return p1_strat(b)
	}
	return p2_strat(b)
}

func p2_strat(b Board) [2]Player {
	x := Player(rand.Intn(3))
	y := Player(rand.Intn(3))
	if b[x][y] == Nil {
		return [2]Player{x, y}
	}
	return p2_strat(b)
}

func assess(b Board, m [2]Player)

func p1_strat(b Board) [2]Player {

}
