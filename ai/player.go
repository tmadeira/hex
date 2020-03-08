/*
Package ai contains a Hex AI.

Usage:

	minDistance, _ := ai.Heuristic("mindistance")
	player := ai.NewPlayer(ai.Max, "minimax", minDistance)
	board := ai.Board{
		Size:   3,
		Matrix: [][]ai.PlayerID{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
	}
	move, outcome, _ := player.Play(board)
*/
package ai

import (
	"fmt"
	"time"
)

// PlayerID identifies a player (NoOne, Max or Min).
type PlayerID int

// Hex players.
const (
	NoOne PlayerID = 0
	Max   PlayerID = 1
	Min   PlayerID = 2
)

// Player contains information about a Hex player, such as its ID,
// strategy and heuristic.
type Player struct {
	ID        PlayerID
	Strategy  string
	Heuristic HeuristicFunc
}

// NewPlayer returns a new AI player with given ID, strategy and heuristic.
func NewPlayer(id PlayerID, strategy string, heuristic HeuristicFunc) *Player {
	return &Player{id, strategy, heuristic}
}

// Play receives a board state and returns a move and the outcome it expects
// from that move.
func (p *Player) Play(b Board) (*Move, int, error) {
	Print(b, false)

	start := time.Now()

	var mv *Move
	var outcome int
	switch p.Strategy {
	case "minimax":
		mv, outcome = p.Minimax(b, 5)
	default:
		return nil, 0, fmt.Errorf("invalid strategy: %s", p.Strategy)
	}

	if mv == nil {
		return nil, 0, fmt.Errorf("no movement")
	}

	elapsed := time.Since(start)

	fmt.Printf("> AI took %s.\n", elapsed)

	b.Matrix[mv.I][mv.J] = p.ID
	Print(b, true)

	return mv, outcome, nil
}
