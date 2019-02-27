import player from './player';
import winner from '../algorithms/winner';

const RED = 1;
const BLUE = 2;

const INFINITY = 1000000;

class minimax extends player {
  constructor(n, heuristic) {
    super(n);
    this.heuristic = heuristic;
  }

  abMinimax(board, whoami) {
    const oponent = whoami === RED ? BLUE : RED;
  
    const rec = (depth, max, alpha, beta) => {
      const turn = max ? whoami : oponent;
      const other = turn === RED ? BLUE : RED;
      if (depth === 0 || winner(board)) {
        return [this.heuristic(board, max ? turn : other)];
      }

      if (max) {
        let value = -INFINITY;
        let move = null;

        for (let x = 0; x < this.n; x++) {
          for (let y = 0; y < this.n; y++) {
            if (!board[x][y]) {
              board[x][y] = turn;
              const [v] = rec(depth - 1, false);
              if (v > value) {
                value = v;
                move = [x, y];
              }
              alpha = Math.max(alpha, value);
              board[x][y] = 0;
              if (alpha >= beta) {
                return [value, move];
              }
            }
          }
        }

        return [value, move];
      }

      let value = INFINITY;
      let move = null;
      for (let x = 0; x < this.n; x++) {
        for (let y = 0; y < this.n; y++) {
          if (!board[x][y]) {
            board[x][y] = turn;
            const [v] = rec(depth - 1, true);
            if (v < value) {
              value = v;
              move = [x, y];
            }
            beta = Math.min(beta, value);
            board[x][y] = 0;
            if (alpha >= beta) {
              return [value, move];
            }
          }
        }
      }

      return [value, move];
    }

    return rec(2, true, -INFINITY, INFINITY);
  }

  play(board, whoami, _last, callback) {
    const [v, m] = this.abMinimax(board, whoami);
    console.log(`v = ${v}`);
    callback(m[0], m[1]);
  }
}

export default minimax;
