package ai

import "testing"

func TestMinDistance(t *testing.T) {
	boardA := &Board{
		Size: 3,
		Matrix: [][]PlayerID{
			{1, 2, 1},
			{0, 0, 0},
			{0, 0, 0},
		},
	}

	boardB := &Board{
		Size: 3,
		Matrix: [][]PlayerID{
			{1, 2, 1},
			{0, 2, 0},
			{0, 0, 0},
		},
	}

	boardC := &Board{
		Size: 3,
		Matrix: [][]PlayerID{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}

	tests := []struct {
		name   string
		board  *Board
		player PlayerID
		want   int
	}{
		{"board_a_max", boardA, Max, 2},
		{"board_a_min", boardA, Min, 2},
		{"board_b_max", boardB, Max, 3},
		{"board_b_min", boardB, Min, 1},
		{"board_c_max", boardC, Max, 3},
		{"board_c_min", boardC, Min, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minDistance(test.board, test.player)
			if got != test.want {
				t.Fatalf("minDistance(%v, %d) = %d; want %d", test.board, test.player, got, test.want)
			}
		})
	}
}

func benchmarkMinDistance(sz int, b *testing.B) {
	board := &Board{
		Size: sz,
	}
	board.Matrix = make([][]PlayerID, sz)
	for i := range board.Matrix {
		board.Matrix[i] = make([]PlayerID, sz)
	}

	for n := 0; n < b.N; n++ {
		minDistance(board, Max)
	}
}

func BenchmarkMinDistance7(b *testing.B) {
	benchmarkMinDistance(7, b)
}

func BenchmarkMinDistance9(b *testing.B) {
	benchmarkMinDistance(9, b)
}

func BenchmarkMinDistance11(b *testing.B) {
	benchmarkMinDistance(11, b)
}
