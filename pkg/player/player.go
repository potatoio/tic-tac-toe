package player

// package main

import (
	"fmt"
)

// Player will be a tic player
type Player struct {
	Wins int
	Team string
}

// Get status of player
func (p Player) Print() {
	fmt.Println("Wins", p.Wins, "Team", p.Team)
}

// increments the wins of the player
func (p *Player) UpdateWins() {
	p.Wins++
}
