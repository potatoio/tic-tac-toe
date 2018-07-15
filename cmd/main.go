package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/potatoio/tic-tac-toe/pkg/board"
	"github.com/potatoio/tic-tac-toe/pkg/player"
)

func cToIndex(x int, y int, length int) int {
	// check that this is a valid move
	if x < 3 && y < 3 && y >= 0 && x >= 0 {
		return x + (y * length)
	}

	return -1
}

func main() {

	togglePlayer := true // player 1 when true and player 2 when false
	matchWon := ""

	// Players
	p1 := player.Player{Team: "x"}
	p2 := player.Player{Team: "o"}

	// Board
	b := board.Board{}
	p2.Print()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Lets play a game")

	// game loop
	for scanner.Scan() {
		input := scanner.Text()
		sPos := strings.Split(input, ",")
		xCoor, _ := strconv.Atoi(sPos[0])
		yCoor, _ := strconv.Atoi(sPos[1])
		index := cToIndex(xCoor, yCoor, 3)

		// Print the board before a move

		if index != -1 && index < len(b) && index >= 0 {
			if togglePlayer {
				b.UpdateBoard(p1.Team, index)
			} else {
				b.UpdateBoard(p2.Team, index)
			}

			matchWon = b.CheckBoard()

			if matchWon == "" {
				togglePlayer = !togglePlayer
				if togglePlayer {
					fmt.Println(p1.Team, " p1 turn")
				} else {
					fmt.Println(p2.Team, " p2 turn")
				}
			} else {
				if matchWon == p1.Team {
					p1.UpdateWins()
					fmt.Println("Player ", p1.Team, "Won")
					p1.Print()
				} else if matchWon == p2.Team {
					p2.UpdateWins()
					fmt.Println("Player ", p2.Team, "Won")
					p2.Print()
				}

				os.Exit(1)
				fmt.Println("Would you like to play a game.... again?")
			}

			b.Print()
		}

	}

}
