package ai

func (p *Player) minimax(depth int, player PlayerID, board *Board) (*Move, int) {
	var oponent = Max
	best := infinity + 1
	if player == Max {
		oponent = Min
		best = -infinity - 1
	}

	// Check if board has a winner.
	winner := Winner(board)
	if winner == Max {
		return nil, infinity
	} else if winner == Min {
		return nil, -infinity
	}

	// No more depth to recurse; return heuristic.
	if depth == 0 {
		v := p.Heuristic(board)
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

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (p *Player) alphaBeta(depth, alpha, beta int, player PlayerID, board *Board) (*Move, int) {
	var oponent = Max
	best := infinity + 1
	if player == Max {
		oponent = Min
		best = -infinity - 1
	}

	// Check if board has a winner.
	winner := Winner(board)
	if winner == Max {
		return nil, infinity
	} else if winner == Min {
		return nil, -infinity
	}

	// No more depth to recurse; return heuristic.
	if depth == 0 {
		v := p.Heuristic(board)
		return nil, v
	}

	possible := possibleMoves(board)
	possible = possible[:20]

	var chosen Move
	for _, current := range possible {
		last := board.LastMove
		board.LastMove = &current
		board.Matrix[current.I][current.J] = player
		_, v := p.alphaBeta(depth-1, alpha, beta, oponent, board)
		board.LastMove = last
		board.Matrix[current.I][current.J] = NoOne

		if (player == Max && v > best) || (player == Min && v < best) {
			best = v
			chosen = current
		}

		if player == Max {
			alpha = max(alpha, v)
		} else {
			beta = min(beta, v)
		}

		if alpha >= beta {
			return &chosen, best
		}
	}

	return &chosen, best
}

// ABMinimax runs minimax with alpha-beta pruning in the given board,
// until the given depth, and returns a move and its expected outcome.
func (p *Player) ABMinimax(board Board, depth int) (*Move, int) {
	return p.alphaBeta(depth, -infinity, infinity, p.ID, &board)
}
