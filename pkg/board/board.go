package board

// package main

import (
	"fmt"
)

// THis is an array to of a 2 x 2 game board
type Board [9]string
type line []int

// print state of the board
func (b Board) Print() {
	fmt.Println(b)
}

// update board
func (b *Board) UpdateBoard(team string, index int) {
	if b[index] == "" {
		b[index] = team
	}
}

// rules for the board
func (b Board) CheckBoard() string {
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

		if b[x1] == b[x2] && b[x0] == b[x1] && b[x0] == b[x2] {
			return b[x0]
		}
	}

	return ""
}

// clears the board
func (b *Board) Reset() {
	b = &Board{}
	fmt.Println(b)
}
