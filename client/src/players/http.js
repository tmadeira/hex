import player from './player';

class http extends player {
  constructor(n, {strategy, depth, heuristic}) {
    super(n);
    this.strategy = strategy;
    this.depth = depth;
    this.heuristic = heuristic;
  }

  play(board, whoami, last, callback) {
    const body = {
      id: whoami,
      strategy: this.strategy,
      depth: this.depth,
      heuristic: this.heuristic,
      size: this.n,
      matrix: board,
      last: last,
    };

    const endpoint = 'http://127.0.0.1:8080/play';
    const request = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    };

    const start = new Date();
    fetch(endpoint, request)
      .then(response => response.json())
      .then(data => {
        const elapsed = new Date() - start;
        console.log('Elapsed:', elapsed, 'ms. Move:', data.move, '. Expected outcome:', data.expectedOutcome, '.');
        callback(data.move[0], data.move[1]);
      });
  }
}

export default http;
