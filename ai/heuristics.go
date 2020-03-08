package ai

import "fmt"

const infinity = 1000000

// HeuristicFunc is a function that receives a board and returns its
// expected outcome for the Max player.
type HeuristicFunc func(*Board) int

// Heuristic returns a heuristic function from a given heuristic name.
func Heuristic(h string) (HeuristicFunc, error) {
	switch h {
	case "mindistance":
		return HeuristicFunc(HeuristicMinDistance), nil
	default:
		return nil, fmt.Errorf("invalid heuristic: %s", h)
	}
}

// HeuristicMinDistance returns the minimum number of moves required for
// Max player to win minus the minimum number of moves required for Min
// player to win.
func HeuristicMinDistance(b *Board) int {
	// TODO.
	return 0
}
