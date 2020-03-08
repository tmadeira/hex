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

func minDistance(b *Board, player PlayerID) int {
	M := make([][]int, b.Size+6)
	for i := range M {
		M[i] = make([]int, b.Size+6)
		for j := range M[i] {
			M[i][j] = -1
		}
	}

	oponent := Max
	if player == Max {
		oponent = Min
	}

	queue := make([]Move, (6+b.Size)*b.Size)
	start := 0
	end := 0

	s := make([]Move, b.Size)
	for i := range s {
		if player == Max {
			s[i] = Move{i, -1}
		} else {
			s[i] = Move{-1, i}
		}
	}
	connected := connect(b, s, player, M, 0)

	for _, v := range connected {
		if inBoard(v, b.Size) || (player == Max && v.J == -1) || (player == Min && v.I == -1) {
			queue[end] = v
			end++
		}
		if (player == Max && v.J == b.Size) || (player == Min && v.I == b.Size) {
			return 0
		}
	}

	for start < end {
		u := queue[start]
		start++

		for d := 0; d < 6; d++ {
			v := Move{u.I + di[d], u.J + dj[d]}
			if M[v.I+3][v.J+3] != -1 {
				continue
			}
			if inBoard(v, b.Size) && b.Matrix[v.I][v.J] != oponent {
				connected := connect(b, []Move{v}, player, M, M[u.I+3][u.J+3]+1)
				for _, v := range connected {
					if inBoard(v, b.Size) {
						queue[end] = v
						end++
					}
					if (player == Max && v.J == b.Size) || (player == Min && v.I == b.Size) {
						return M[u.I+3][u.J+3] + 1
					}
				}
			}
		}
	}

	return infinity
}

// HeuristicMinDistance returns the minimum number of moves required for
// Max player to win minus the minimum number of moves required for Min
// player to win.
func HeuristicMinDistance(b *Board) int {
	return minDistance(b, Max) - minDistance(b, Min)
}
