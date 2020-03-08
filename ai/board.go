package ai

import "fmt"

var (
	di = []int{0, 1, 1, 0, -1, -1}
	dj = []int{-1, -1, 0, 1, 1, 0}
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

func connect(b *Board, start []Move, id PlayerID, M [][]int, dist int) []Move {
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
	}

	for i := 0; i < len(start); i++ {
		M[start[i].I+3][start[i].J+3] = dist
		connected = append(connected, start[i])
		dfs(start[i])
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
	connect(b, start, id, M, 0)

	for i := 0; i < len(end); i++ {
		if M[end[i].I+3][end[i].J+3] != -1 {
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
