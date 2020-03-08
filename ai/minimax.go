package ai

func (p *Player) minimax(depth int, player PlayerID, board *Board) (*Move, int) {
	var oponent = Max
	best := infinity
	if player == Max {
		oponent = Min
		best = -infinity
	}

	// Check if board has a winner.
	winner := Winner(board)
	if winner == player {
		return nil, infinity
	} else if winner == oponent {
		return nil, -infinity
	}

	// No more depth to recurse; return heuristic.
	if depth == 0 {
		v := p.Heuristic(board)
		if player == Min {
			v = -v
		}

		return nil, v
	}

	var mv *Move
	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			if board.Matrix[i][j] == NoOne {
				board.Matrix[i][j] = player
				_, v := p.minimax(depth-1, oponent, board)
				board.Matrix[i][j] = NoOne

				if player == Max && v > best {
					best = v
					mv = &Move{i, j}
				} else if player == Min && v < best {
					best = v
					mv = &Move{i, j}
				}
			}
		}
	}

	return mv, best
}

// Minimax runs minimax in the given board, until the given depth, and
// returns a move and its expected outcome.
func (p *Player) Minimax(board Board, depth int) (*Move, int) {
	return p.minimax(depth, p.ID, &board)
}
