package main

import (
	"math/rand"
	"sort"
)

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

func assess(b Board, m [2]Player) Assessed {
	bc := Place(b, 1, m)
	if Win(bc, 1) {
		return Assessed{move: m, score: 100}
	}
	return Assessed{move: m, score: 50}

}

func p1_strat(b Board) [2]Player {
	var assessed []Assessed
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == Nil {
				assessed = append(assessed, assess(b, [2]Player{Player(i), Player(j)}))
			}
		}
	}
	sort.Slice(assessed, func(i, j int) bool { return assessed[i].score > assessed[j].score })
	move := assessed[0].move
	return move
}
