package main

import (
	"math/rand"
	"sort"
)

type Assessed struct {
	ev   int
	move [2]Player
}

func Simulate(b Board, p Player, n int) [3]int {
	// Returns the outcomes of playing n games
	// results[0] = draws
	// results[1] = p1 wins
	// results[2] = p2 wins
	results := [3]int{0, 0, 0}
	for i := 0; i < n; i++ {
		results[Run(b, p)]++
	}
	return results
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

func Assess(b Board, m [2]Player) Assessed {
	bc := Place(b, 1, m)
	if Win(bc, 1) {
		return Assessed{move: m, ev: 100}
	}
	results := Simulate(bc, 2, 10)
	winrate := (results[1] / 10) * 100
	return Assessed{move: m, ev: winrate}
}

func p1_strat(b Board) [2]Player {
	var assessed []Assessed
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == Nil {
				assessed = append(assessed, Assess(b, [2]Player{Player(i), Player(j)}))
			}
		}
	}
	sort.Slice(assessed, func(i, j int) bool { return assessed[i].ev > assessed[j].ev })
	move := assessed[0].move
	return move
}
