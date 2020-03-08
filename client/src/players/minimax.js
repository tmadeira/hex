import player from './player';
import winner from '../algorithms/winner';

const RED = 1;
const BLUE = 2;

const INFINITY = 1000000;
const MAX_TESTS = 20;

const dx = [0, 1, 1, 0, -1, -1];
const dy = [-1, -1, 0, 1, 1, 0];

class minimax extends player {
  constructor(n, heuristic) {
    super(n);
    this.heuristic = heuristic;
  }

  abMinimax(board, whoami, last) {
    const oponent = whoami === RED ? BLUE : RED;
  
    const rec = (depth, max, alpha, beta, last) => {
      const turn = max ? whoami : oponent;
      const other = turn === RED ? BLUE : RED;
      if (depth === 0) {
        return [this.heuristic(board, max ? turn : other)];
      }

      if (winner(board)) {
        return [2 * this.heuristic(board, max ? turn : other)];
      }

      if (max) {
        let value = -INFINITY;
        let move = null;

        const M = {};
        const queue = [];
        queue.push(last);
        M[`${last[0]} ${last[1]}`] = true;
        let tested = 0;

        while (queue.length && tested < MAX_TESTS) {
          const [x, y] = queue.shift();
          if (!board[x][y]) {
            board[x][y] = turn;
            const [v] = rec(depth - 1, false, alpha, beta, [x, y]);
            if (v > value) {
              value = v;
              move = [x, y];
            }
            alpha = Math.max(alpha, value);
            board[x][y] = 0;
            if (alpha >= beta) {
              return [value, move];
            }
            tested++;
          }

          for (let d = 0; d < dx.length; d++) {
            const nx = x + dx[d];
            const ny = y + dy[d];
            if (nx < 0 || ny < 0 || nx >= this.n || ny >= this.n) {
              continue;
            }

            const np = `${nx} ${ny}`;
            if (np in M) {
              continue;
            }

            M[np] = true;
            queue.push([nx, ny]);
          }
        }

        return [value, move];
      }

      let value = INFINITY;
      let move = null;

      const M = {};
      const queue = [];
      queue.push(last);
      M[`${last[0]} ${last[1]}`] = true;
      let tested = 0;

      while (queue.length && tested < MAX_TESTS) {
        const [x, y] = queue.shift();
        if (!board[x][y]) {
          board[x][y] = turn;
          const [v] = rec(depth - 1, true, alpha, beta, [x, y]);
          if (v < value) {
            value = v;
            move = [x, y];
          }
          beta = Math.min(beta, value);
          board[x][y] = 0;
          if (alpha >= beta) {
            return [value, move];
          }
          tested++;
        }

        for (let d = 0; d < dx.length; d++) {
          const nx = x + dx[d];
          const ny = y + dy[d];
          if (nx < 0 || ny < 0 || nx >= this.n || ny >= this.n) {
            continue;
          }

          const np = `${nx} ${ny}`;
          if (np in M) {
            continue;
          }

          M[np] = true;
          queue.push([nx, ny]);
        }
      }

      return [value, move];
    }

    return rec(4, true, -INFINITY, INFINITY, last ? last : [Math.floor(this.n/2), Math.floor(this.n/2)]);
  }

  play(board, whoami, last, callback) {
    const [v, m] = this.abMinimax(board, whoami, last);
    console.log(`v = ${v}`);
    callback(m[0], m[1]);
  }
}

export default minimax;
