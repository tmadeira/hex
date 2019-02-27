const dx = [0, 1, 1, 0, -1, -1];
const dy = [-1, -1, 0, 1, 1, 0];

const RED = 1;
const BLUE = 2;

const bfs = (board, start, player) => {
  const n = board.length;

  const D = {};
  const queue = [];

  for (let i = 0; i < start.length; i++) {
    const [x, y] = start[i];
    queue.push([x, y]);
    D[`${x} ${y}`] = 0;

    const connectedQueue = [];
    connectedQueue.push([x, y]);
    while (connectedQueue.length > 0) {
      const [ox, oy] = connectedQueue.shift();

      for (let e = 0; e < dx.length; e++) {
        const px = ox + dx[e];
        const py = oy + dy[e];

        if (px < 0 || py < 0 || px >= n || py >= n) {
          continue;
        }

        if (board[px][py] !== player) {
          continue;
        }

        const ppair = `${px} ${py}`;
        if (ppair in D) {
          continue;
        }

        D[`${px} ${py}`] = 0;
        queue.push([px, py]);
        connectedQueue.push([px, py]);
      }
    }
  }

  while (queue.length > 0) {
    const [x, y] = queue.shift();
    const pair = `${x} ${y}`;

    for (let d = 0; d < dx.length; d++) {
      const nx = x + dx[d];
      const ny = y + dy[d];
      if (nx < 0 || ny < 0 || nx >= n || ny >= n) {
        continue;
      }

      if (board[nx][ny]) {
        continue;
      }

      const npair = `${nx} ${ny}`;
      if (npair in D) {
        continue;
      }

      const connectedQueue = [];

      D[`${nx} ${ny}`] = D[pair] + 1;
      queue.push([nx, ny]);
      connectedQueue.push([nx, ny]);
      while (connectedQueue.length > 0) {
        const [ox, oy] = connectedQueue.shift();

        for (let e = 0; e < dx.length; e++) {
          const px = ox + dx[e];
          const py = oy + dy[e];

          if (px < 0 || py < 0 || px >= n || py >= n) {
            continue;
          }

          if (board[px][py] !== player) {
            continue;
          }

          const ppair = `${px} ${py}`;
          if (ppair in D) {
            continue;
          }

          D[`${px} ${py}`] = D[pair] + 1;
          queue.push([px, py]);
          connectedQueue.push([px, py]);
        }
      }
    }
  }

  return D;
};

const smallest = (board) => {
  const n = board.length;
  const red = [];
  const blue = [];

  for (let i = 0; i < n; i++) {
    red.push([i, -1]);
    blue.push([-1, i]);
  }

  const rd = bfs(board, red, RED);
  const bd = bfs(board, blue, BLUE);

  let rs = n*n;
  let bs = n*n;
  for (let i = 0; i < n; i++) {
    const rp = `${i} ${n-1}`;
    const bp = `${n-1} ${i}`;
    if (rp in rd) {
      rs = Math.min(rs, rd[rp]);
    }
    if (bp in bd) {
      bs = Math.min(bs, bd[bp]);
    }
  }

  return [rs, bs];
};

const heuristic = (board, whoami) => {
  const [red, blue] = smallest(board);
  return whoami === 1 ? blue - red : red - blue;
};

export default smallest;
export {heuristic};
