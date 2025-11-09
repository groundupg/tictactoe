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
	n := 1
	results := [3]int{0, 0, 0}
	for i := 0; i < n; i++ {
		b := Init()
		first := flip()
		results[Run(b, first)]++
	}

	fmt.Print("GAMES PLAYED: ", n, "\r\n")
	fmt.Print("P1 WIN RATE: ", float64(results[1])/float64(n)*100, "%\r\n")
	fmt.Print("P2 WIN RATE: ", float64(results[2])/float64(n)*100, "%\r\n")
	fmt.Print("DRAW RATE: ", float64(results[0])/float64(n)*100, "%\r\n")
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

func Place(b Board, p Player, i [2]Player) Board {
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

func flip() Player {
	return Player(rand.Intn(2) + 1)
}
