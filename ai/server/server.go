package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tmadeira/hex/ai"
)

func handlePlay() http.HandlerFunc {
	type request struct {
		PlayerID  int     `json:"id"`
		Strategy  string  `json:"strategy"`
		Depth     int     `json:"depth"`
		Heuristic string  `json:"heuristic"`
		Size      int     `json:"size"`
		Matrix    [][]int `json:"matrix"`
		Last      []int   `json:"last"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		in := request{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&in)
		if err != nil {
			http.Error(w, "Can't decode body", http.StatusBadRequest)
			return
		}

		b := ai.Board{}

		b.Size = in.Size

		b.Matrix = make([][]ai.PlayerID, in.Size)
		for i := range b.Matrix {
			b.Matrix[i] = make([]ai.PlayerID, in.Size)
		}
		for i := 0; i < b.Size; i++ {
			for j := 0; j < b.Size; j++ {
				b.Matrix[i][j] = ai.PlayerID(in.Matrix[i][j])
			}
		}

		if len(in.Last) == 2 {
			b.LastMove = &ai.Move{I: in.Last[0], J: in.Last[1]}
		}

		heuristicFunc, err := ai.Heuristic(in.Heuristic)
		if err != nil {
			http.Error(w, "Invalid heuristic "+in.Heuristic, http.StatusBadRequest)
			return
		}

		player := ai.NewPlayer(ai.PlayerID(in.PlayerID), in.Strategy, in.Depth, heuristicFunc)
		mv, v, err := player.Play(b)
		if err != nil {
			http.Error(w, fmt.Sprintf("Can't play: %v", err), http.StatusInternalServerError)
			return
		}

		move := make([]int, 0)
		if mv != nil {
			move = []int{mv.I, mv.J}
		}

		js, err := json.Marshal(map[string]interface{}{
			"move":            move,
			"expectedOutcome": v,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Can't marshal JSON: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

// Run runs Hex AI HTTP server in the specified TCP port.
func Run(port int) {
	http.HandleFunc("/play", handlePlay())

	log.Printf("Starting server on port %d", port)
	listen := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listen, nil)
	log.Fatal(err)
}
