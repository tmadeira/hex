import player from './player';

class http extends player {
  constructor(n) {
    super(n);
  }

  play(board, whoami, last, callback) {
    const body = {
      size: this.n,
      matrix: board,
      last: last,
    };

    console.log('body', body);

    const endpoint = 'http://127.0.0.1:8080/play';
    const request = {
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    };

    fetch(endpoint, request)
      .then(response => { console.log(response); });

    // TODO: Make move.
  }
}

export default minimax;
