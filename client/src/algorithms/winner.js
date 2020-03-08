const RED = 1;
const BLUE = 2;

const dx = [0, 1, 1, 0, -1, -1];
const dy = [-1, -1, 0, 1, 1, 0];

const winner = (board) => {
  let M = {};
  const n = board.length - 1;

  const dfs = (who, x, y) => {
    if (who === RED && y === n - 1) {
      return true;
    }
    if (who === BLUE && x === n - 1) {
      return true;
    }

    const pair = `${x} ${y}`;
    M[pair] = true;
    for (let d = 0; d < dx.length; d++) {
      const nx = x + dx[d];
      const ny = y + dy[d];
      if (nx < 0 || ny < 0 || nx >= n || ny >= n) {
        continue;
      }

      const npair = `${nx} ${ny}`;
      if (!(npair in M) && board[nx][ny] === who) {
        if (dfs(who, nx, ny)) {
          return true;
        }
      }
    }
    return false;
  };

  for (let i = 0; i < n; i++) {
    if (board[i][0] === RED) {
      M = {};
      if (dfs(RED, i, 0)) {
        return RED;
      }
    }
    if (board[0][i] === BLUE) {
      M = {};
      if (dfs(BLUE, 0, i)) {
        return BLUE;
      }
    }
  }
  return 0;
}

export default winner;
