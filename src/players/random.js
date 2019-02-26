import player from './player';

class random extends player {
  play(board, _whoami, _last, callback) {
    const possible = [];
    for (let x = 0; x < this.n; x++) {
      for (let y = 0; y < this.n; y++) {
        if (!board[x][y]) {
          possible.push([x, y]);
        }
      }
    }

    const chosen = possible[Math.floor(Math.random() * possible.length)];
    callback(chosen[0], chosen[1]);
  }
}

export default random;
