class player {
  constructor(n) {
    this.n = n;
  }

  randomMove(board) {
    const possible = [];
    for (let x = 0; x < this.n; x++) {
      for (let y = 0; y < this.n; y++) {
        if (!board[x][y]) {
          possible.push([x, y]);
        }
      }
    }

    const chosen = possible[Math.floor(Math.random() * possible.length)];
    return chosen;
  }
}

export default player;
