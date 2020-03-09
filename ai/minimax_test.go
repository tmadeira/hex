package ai

import (
	"math/rand"
	"testing"
)

func benchmarkMinimax(sz, depth int, ab bool, b *testing.B) {
	rand.Seed(1)
	p := &Player{
		ID:        Max,
		Heuristic: HeuristicMinDistance,
	}

	for n := 0; n < b.N; n++ {
		board := &Board{
			Size: sz,
		}
		board.Matrix = make([][]PlayerID, sz)
		for i := range board.Matrix {
			board.Matrix[i] = make([]PlayerID, sz)
		}

		if ab {
			p.ABMinimax(*board, depth)
		} else {
			p.Minimax(*board, depth)
		}
	}
}

func BenchmarkMinimax4_4(b *testing.B) {
	benchmarkMinimax(4, 4, false, b)
}

func BenchmarkABMinimax4_4(b *testing.B) {
	benchmarkMinimax(4, 4, true, b)
}

func BenchmarkABMinimax4_6(b *testing.B) {
	benchmarkMinimax(4, 6, true, b)
}

func BenchmarkABMinimax9_4(b *testing.B) {
	benchmarkMinimax(9, 4, true, b)
}
