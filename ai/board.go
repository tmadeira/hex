package ai

import "fmt"

// Neighbors.
var (
	di = []int{0, 1, 1, 0, -1, -1}
	dj = []int{-1, -1, 0, 1, 1, 0}
)

// Bridges.
var (
	bi = []int{1, 2, 1, -1, -2, -1}
	bj = []int{-2, -1, 1, 2, 1, -1}
)

// Board stores a hex board state.
type Board struct {
	Size     int
	Matrix   [][]PlayerID
	LastMove *Move
}

func inBoard(m Move, size int) bool {
	if m.I < 0 || m.J < 0 {
		return false
	}
	if m.I >= size || m.J >= size {
		return false
	}
	return true
}

func connect(b *Board, start []Move, id PlayerID, M [][]int, dist int, bridges bool) []Move {
	var connected []Move

	var dfs func(u Move)
	dfs = func(u Move) {
		for d := 0; d < 6; d++ {
			v := Move{u.I + di[d], u.J + dj[d]}
			if M[v.I+3][v.J+3] != -1 {
				continue
			}

			if inBoard(v, b.Size) && b.Matrix[v.I][v.J] == id {
				M[v.I+3][v.J+3] = dist
				connected = append(connected, v)
				dfs(v)
			} else if !inBoard(v, b.Size) {
				M[v.I+3][v.J+3] = dist
				connected = append(connected, v)
			}
		}

		if bridges {
			for d := 0; d < 6; d++ {
				v := Move{u.I + bi[d], u.J + bj[d]}
				if M[v.I+3][v.J+3] != -1 {
					continue
				}

				x := Move{u.I + di[d], u.J + dj[d]}
				y := Move{u.I + di[(d+1)%6], u.J + dj[(d+1)%6]}
				if !inBoard(x, b.Size) || !inBoard(y, b.Size) {
					continue
				}

				if b.Matrix[x.I][x.J] != NoOne || b.Matrix[y.I][y.J] != NoOne {
					continue
				}

				if inBoard(v, b.Size) && b.Matrix[v.I][v.J] == id {
					M[v.I+3][v.J+3] = dist
					connected = append(connected, v)
					dfs(v)
				} else if !inBoard(v, b.Size) {
					M[v.I+3][v.J+3] = dist
					connected = append(connected, v)
				}
			}
		}
	}

	for _, s := range start {
		M[s.I+3][s.J+3] = dist
		connected = append(connected, s)
		dfs(s)
	}

	return connected
}

func connected(b *Board, start, end []Move, id PlayerID) bool {
	M := make([][]int, b.Size+6)
	for i := range M {
		M[i] = make([]int, b.Size+6)
		for j := range M[i] {
			M[i][j] = -1
		}
	}
	connect(b, start, id, M, 0, false)

	for _, e := range end {
		if M[e.I+3][e.J+3] != -1 {
			return true
		}
	}

	return false
}

// Print prints the given board to standard output.
func Print(b Board, indent bool) {
	for i := 0; i < b.Size; i++ {
		if indent {
			fmt.Printf("> ")
		}
		for j := 0; j < i; j++ {
			fmt.Printf("  ")
		}
		for j := 0; j < b.Size; j++ {
			fmt.Printf(" %d", b.Matrix[j][i])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// Winner checks if there is a winner in a given board state. It returns
// the winner's ID or NoOne in case there is no winner.
func Winner(b *Board) PlayerID {
	start := make([]Move, b.Size)
	end := make([]Move, b.Size)

	for i := 0; i < b.Size; i++ {
		start[i] = Move{i, -1}
		end[i] = Move{i, b.Size}
	}
	if connected(b, start, end, Max) {
		return Max
	}

	for i := 0; i < b.Size; i++ {
		start[i] = Move{-1, i}
		end[i] = Move{b.Size, i}
	}
	if connected(b, start, end, Min) {
		return Min
	}

	return NoOne
}

// possibleMoves return the possible moves in the given board, ordering
// moves by distance from the last move.
func possibleMoves(b *Board) []Move {
	var possible []Move

	M := make([][]bool, b.Size)
	for i := range M {
		M[i] = make([]bool, b.Size)
	}

	queue := make([]Move, b.Size*b.Size)
	start := 0
	end := 1
	if b.LastMove != nil {
		queue[0] = *b.LastMove
		M[b.LastMove.I][b.LastMove.J] = true
	} else {
		queue[0] = Move{0, 0}
		M[0][0] = true
	}

	for start < end {
		u := queue[start]
		start++
		for d := 0; d < 6; d++ {
			v := Move{u.I + di[d], u.J + dj[d]}
			if !inBoard(v, b.Size) {
				continue
			}
			if M[v.I][v.J] {
				continue
			}
			M[v.I][v.J] = true
			if b.Matrix[v.I][v.J] == NoOne {
				possible = append(possible, v)
			}
			queue[end] = v
			end++
		}
	}

	return possible
}
