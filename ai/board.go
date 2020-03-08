package ai

// Board stores a hex board state.
type Board struct {
	Size     int
	Matrix   [][]PlayerID
	LastMove *Move
}

// Winner checks if there is a winner in a given board state. It returns
// the winner's ID or NoOne in case there is no winner.
func Winner(b *Board) PlayerID {
	return NoOne
}
