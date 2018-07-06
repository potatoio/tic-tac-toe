package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type player struct {
	wins   int
	symbol string
}

type board [9]string
type line []int

func (p player) print() {
	fmt.Println("Wins", p.wins, "Symbol", p.symbol)
}

func (p *player) updateWins() {
	p.wins = p.wins + 1
}

func (b board) print() {
	fmt.Println(b)
}

func cToIndex(x int, y int, length int) int {
	// check that this is a valid move
	return x + (y * length)
}

func (b *board) updateBoard(symbol string, index int) {
	if b[index] == "" {
		b[index] = symbol
	}

	fmt.Println("Invalid position")
}

func (b board) checkBoard() string {
	fmt.Println("validate board")
	lines := [][]int{
		line{0, 1, 2},
		line{3, 4, 5},
		line{6, 7, 8},
		line{0, 3, 6},
		line{1, 4, 7},
		line{2, 5, 8},
		line{0, 4, 8},
		line{2, 4, 6},
	}

	for _, l := range lines {
		x0, x1, x2 := l[0], l[1], l[2]

		// fmt.Println(a, b, c)
		if b[x1] == b[x2] && b[x0] == b[x1] && b[x0] == b[x2] {
			return b[x0]
		}

	}

	return ""
}

func main() {

	togglePlayer := true // player 1 when true and player 2 when false
	matchWon := ""

	// Players
	p1 := player{symbol: "x"}
	p2 := player{symbol: "o"}

	// board
	b := board{}

	p1.print()
	p2.print()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Lets play a game")

	// game loop
	for scanner.Scan() {
		input := scanner.Text()
		sPos := strings.Split(input, ",")
		xCoor, _ := strconv.Atoi(sPos[0])
		yCoor, _ := strconv.Atoi(sPos[1])
		index := cToIndex(xCoor, yCoor, 3)
		fmt.Println(index)

		// Print the board before a move
		b.print()

		if togglePlayer {
			b.updateBoard(p1.symbol, index)
		} else {
			b.updateBoard(p2.symbol, index)
		}

		matchWon = b.checkBoard()

		if matchWon == "" {
			fmt.Println(matchWon, "nothing yet")
		} else {
			if matchWon == p1.symbol {
				p1.updateWins()
				p1.print()
			} else if matchWon == p2.symbol {
				p2.updateWins()
				p2.print()
			}

			// clean the board here
			// ask if player wants to win play again.
		}

		// Print the board after a move
		b.print()

		togglePlayer = !togglePlayer

	}

}
