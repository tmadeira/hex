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

	boardD := &Board{
		Size: 5,
		Matrix: [][]PlayerID{
			{0, 0, 0, 2, 0},
			{0, 2, 1, 0, 1},
			{0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	}

	boardE := &Board{
		Size: 5,
		Matrix: [][]PlayerID{
			{0, 2, 0, 0, 0},
			{2, 1, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	}

	boardF := &Board{
		Size: 5,
		Matrix: [][]PlayerID{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 2, 1, 0},
			{0, 1, 0, 0, 0},
			{0, 2, 0, 0, 0},
		},
	}

	boardG := &Board{
		Size: 5,
		Matrix: [][]PlayerID{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 2, 1, 0},
			{0, 1, 1, 0, 0},
			{0, 2, 0, 0, 0},
		},
	}

	boardH := &Board{
		Size: 7,
		Matrix: [][]PlayerID{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 2, 0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0},
			{0, 2, 0, 0, 0, 0, 0},
		},
	}

	tests := []struct {
		name    string
		board   *Board
		player  PlayerID
		bridges bool
		want    int
	}{
		{"board_a_max", boardA, Max, false, 2},
		{"board_a_min", boardA, Min, false, 2},
		{"board_b_max", boardB, Max, false, 3},
		{"board_b_min", boardB, Min, false, 1},
		{"board_c_max", boardC, Max, false, 3},
		{"board_c_min", boardC, Min, false, 3},
		{"board_d_max", boardD, Max, false, 2},
		{"board_d_min", boardD, Min, false, 4},
		{"board_e_max", boardE, Max, false, 3},
		{"board_e_min", boardE, Min, false, 5},
		{"board_f_max", boardF, Max, true, 1},
		{"board_f_min", boardF, Min, true, 2},
		{"board_g_max", boardG, Max, true, 0},
		{"board_g_min", boardG, Min, true, 3},
		{"board_h_max", boardH, Max, true, 1},
		{"board_h_min", boardH, Min, true, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minDistance(test.board, test.player, test.bridges, false)
			if got != test.want {
				t.Fatalf("minDistance(%v, %d) = %d; want %d", test.board, test.player, got, test.want)
			}
		})
	}
}

func benchmarkMinDistance(sz int, bridges, both bool, b *testing.B) {
	board := &Board{
		Size: sz,
	}
	board.Matrix = make([][]PlayerID, sz)
	for i := range board.Matrix {
		board.Matrix[i] = make([]PlayerID, sz)
	}

	for n := 0; n < b.N; n++ {
		minDistance(board, Max, bridges, both)
	}
}

func BenchmarkMinDistance7(b *testing.B) {
	benchmarkMinDistance(7, false, false, b)
}

func BenchmarkMinDistance9(b *testing.B) {
	benchmarkMinDistance(9, false, false, b)
}

func BenchmarkMinDistance11(b *testing.B) {
	benchmarkMinDistance(11, false, false, b)
}

func BenchmarkMinDistanceBridges7(b *testing.B) {
	benchmarkMinDistance(7, true, false, b)
}

func BenchmarkMinDistanceBridges9(b *testing.B) {
	benchmarkMinDistance(9, true, false, b)
}

func BenchmarkMinDistanceBridges11(b *testing.B) {
	benchmarkMinDistance(11, true, false, b)
}

func BenchmarkMinDistanceBridgesBoth7(b *testing.B) {
	benchmarkMinDistance(7, true, true, b)
}

func BenchmarkMinDistanceBridgesBoth9(b *testing.B) {
	benchmarkMinDistance(9, true, true, b)
}

func BenchmarkMinDistanceBridgesBoth11(b *testing.B) {
	benchmarkMinDistance(11, true, true, b)
}
