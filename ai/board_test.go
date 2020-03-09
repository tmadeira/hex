package ai

import (
	"math/rand"
	"testing"
)

func randomPlayer() PlayerID {
	return PlayerID(rand.Intn(4) % 3)
}

func benchmarkWinner(sz int, b *testing.B) {
	rand.Seed(1)

	for n := 0; n < b.N; n++ {
		board := &Board{
			Size: sz,
		}
		board.Matrix = make([][]PlayerID, sz)
		for i := range board.Matrix {
			board.Matrix[i] = make([]PlayerID, sz)
		}

		Winner(board)
	}
}

func BenchmarkWinner7(b *testing.B) {
	benchmarkWinner(7, b)
}

func BenchmarkWinner9(b *testing.B) {
	benchmarkWinner(9, b)
}

func BenchmarkWinner11(b *testing.B) {
	benchmarkWinner(11, b)
}
