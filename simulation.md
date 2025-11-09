# Simulation

The primary utility of a computer is it's capacity for mass computation. It is this capacity for
large computation which saw the computer replacing humans as the defacto means of computation.
This utility remains the primary use of the computer today.

What does not require large computation may not require a computer; it may be argued that the task of defining a program which must be *computable* on a machine is a layer which, if possible, one
should avoid.

However, we should not disregard the computer; the utility of large computation is in fact one very
useful -- one which may enable the human to reach heights never before achievable.

What I propose is the utilisation of the computer's ability to perform large computations.

## On Decision Making

*A decision can be weighted by the expected value of making the decision.*

*The expected value of a decision is the weighted sum of the multiplication of the value and
probability of all possible outcomes**

```
  type Outcome:
      probability: float
      value: float

  fn o_ev (p: float v: float) -> float:
      return p * v

  fn EV (no: list[Outcome]) -> float:
      current: float
      for o in no:
        current += o_ev(o.probability, o.value)
      return current 
```

The above are very simple computations to calculate the EV of a given decision. However, the computation
has a requirement -- the list of possible outcomes of making a decision, with their probability
& value.

It is precisely these areas where humans struggle, in establishing the possible outcomes of taking
a decision. Very simple games can be handled, i.e. checkers, but when more factors are introduced
humans crumble.

I propose that the most valuable way we, as humans, can utilise the computer, is through simulation.
Specifically, simulation of possible outcomes of a decision which enable a human to make expected
value decisions.

This has proven to be extremely beneficial in many cases: chess, poker, go. The afore mentioned
are games with clear constraints that can be defined, enabling the mass simulation of outcomes
which facilitate expected value decisions for any single decision.

I, however, think that we can do more.
I think that we can use the simulation capacity of computers to apply these methods to business &
science.

## A Simulator

If we are to take the *Monte Carlo Method* as a means of building an initial simulator, we may
then have a view on how a simulator program may be built for a specific use case.

The Monte Carlo Method takes a stepped pattern, using random sampling based on a probability
distribution to perform a deterministic computation.

The method takes a stepped pattern:

    1. Define a domain of possible inputs
    2. Generate inputs randomly from a probability distribution over the domain
    3. Perform a deterministic computation of the outputs
    4. Aggregate the results

We will, in the following section, take this method as a reference in building a simulator program
for our use case(s).

### Tic Tac Toe

*statements*

tic tac toe is a turn-based game.
tic tac toe has 2 players.
player 1 has a mark of 1, player 2 has a mark of 2.
tic tac toe is played on a 3x3 board.
each cell on the board can either be empty, or one of the player's marks.
a win occurs when a player has 3 joined cells -- diagonal, horizontal, or veritcal.
on every turn a player places their mark on an empty cell.
if all cells are filled and there is no win, the game ends in a draw.

*types*

type Cell: enum{0, 1, 2} # Empty, p1, or p2
type Board: [3][3]Cell
type Outcome: enum{0, 1, 2} # Draw, Player 1 win, Player 2 Win
type Player: enum{1, 2} # Player 1 or 2's turn

*functions*

fn init_b() -> Board:
    return [[0, 0, 0], [0, 0, 0], [0, 0, 0]]

fn place(p: Player, b: Board, i) -> Board:
    return b[i] = p

fn swap(t: Turn):
    if t == 1:
        return 2
    return 1

fn run(b: Board, p: Player) -> Outcome:
    b = place(p, b, determine_i(b, p))
    if win(b):
        return p
    if draw(b):
        return 0
    return run(b, swap(t))

fn win(b: Board) -> bool:
    
    
*main event*

run(init_b(), randomise())

## 09.11.25 -- Implementation of a Tic Tac Toe Simulator

I have took the above as a rough guide in implementing a simulator for Tic Tac Toe, in the programming
language golang.

The types used:
```go
type Player uint8
type Board [3][3]Player

const (
	Nil Player = iota
	P1
	P2
)
```

The main event:
```go
func main() {
	b := Init()
	first := flip()
	fmt.Print("FIRST: ", first, "\n")
	fmt.Print(b, "\n")
	Run(b, first)
}

func Init() Board {
	return Board{{Nil, Nil, Nil}, {Nil, Nil, Nil}, {Nil, Nil, Nil}}
}

func Run(b Board, p Player) Player {
	b = Place(b, p, determine_i(b, p))
	fmt.Print(b, "\n")
	if Win(b, p) {
		fmt.Printf("WINNER: %d\n", p)
		return p
	}
	if Draw(b) {
		fmt.Printf("DRAW\n")
		return 0
	}
	return Run(b, Swap(p))
}

func Place(b Board, p Player, i [2]int) Board {
	bc := b
	bc[i[0]][i[1]] = p
	return bc
}
```

The `main()` function simulates the execution of a single game of tic tac toe between two players.
It is clear from looking at the execution of the `main()` function that the process of the game
is detailed in the runtime execution of the `Run` function.

If we are to examine these functions from the view of examining both player's strategy it becomes
obvious that the measure of strategy of a given player is dependent upon the output of the function `determine_i(b, p)`.
The `Place` function uses the output of `determine_i` to make a move.

Currently, the state of `determine_i` produces a random legal move on the board. There is no strategic
logic embedded within this function, resulting in a random winner each time.

Let us change this; instead of a random winner each time, we will apply tic tac toe strategy to
one of the players, with the hope they will develop a greater win rate than the other.

However, before we apply this, let us bring the program to a state where we can measure the win
rate of each. Let us prove win rate.

### Proving Win Rate

We can define *win rate* as the % of wins over n games played.
The execution of `Run` produces an outcome of 0 for draw, 1 for P1 win, 2 for P2 win.


```go
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
```

Running this computation confirms the Win Rate for P1 & P2 are effectively the same:
```shell
go run simulation.go

# Standard Output
P1 WIN RATE: 43.6457
P2 WIN RATE: 43.65547
DRAW RATE: 12.69883
```

