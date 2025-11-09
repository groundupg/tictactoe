package main

import "math/rand"

import "fmt"

type Player uint8
type Board [3][3]Player

const (
	Nil Player = iota
	P1
	P2
)

func main() {
	n := 10000000
	results := [3]int{0, 0, 0}
	for i := 0; i < n; i++ {
		b := Init()
		first := flip()
		results[Run(b, first)]++
	}

	fmt.Print("P1 WIN RATE: ", float64(results[1])/float64(n)*100, "\r\n")
	fmt.Print("P2 WIN RATE: ", float64(results[2])/float64(n)*100, "\r\n")
	fmt.Print("DRAW RATE: ", float64(results[0])/float64(n)*100, "\r\n")
}

func Init() Board {
	return Board{{Nil, Nil, Nil}, {Nil, Nil, Nil}, {Nil, Nil, Nil}}
}

func Run(b Board, p Player) Player {
	b = Place(b, p, DetermineMove(b, p))
	if Win(b, p) {
		return p
	}
	if Draw(b) {
		return 0
	}
	return Run(b, Swap(p))
}

func Place(b Board, p Player, i [2]int) Board {
	bc := b
	bc[i[0]][i[1]] = p
	return bc
}

func Swap(p Player) Player {
	if p == 1 {
		return 2
	}
	return 1
}

func Win(b Board, p Player) bool {
	if horizontal_win(b, p) {
		return true
	}
	if vertical_win(b, p) {
		return true
	}
	if diagonal_win(b, p) {
		return true
	}
	return false
}

func Draw(b Board) bool {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func horizontal_win(b Board, p Player) bool {
	for i := 0; i < len(b); i++ {
		if p == b[i][0] {
			if p == b[i][1] {
				if p == b[i][2] {
					return true
				}
			}
		}
	}
	return false
}

func diagonal_win(b Board, p Player) bool {
	if p == b[1][1] {
		if p == b[0][0] {
			if p == b[2][2] {
				return true
			}
		}
		if p == b[0][2] {
			if p == b[2][0] {
				return true
			}
		}
	}
	return false
}

func vertical_win(b Board, p Player) bool {
	for i := 0; i < len(b); i++ {
		if p == b[0][i] {
			if p == b[1][i] {
				if p == b[2][i] {
					return true
				}
			}
		}
	}
	return false
}

func DetermineMove(b Board, p Player) [2]int {
	x := rand.Intn(3)
	y := rand.Intn(3)
	if b[x][y] == Nil {
		return [2]int{x, y}
	}
	return DetermineMove(b, p)
}

func flip() Player {
	return Player(rand.Intn(2) + 1)
}
