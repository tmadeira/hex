package ai

import "fmt"

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
	switch p.Strategy {
	case "minimax":
		mv, v := p.Minimax(b, 4)
		return mv, v, nil
	default:
		return nil, 0, fmt.Errorf("invalid strategy: %s", p.Strategy)
	}
}
