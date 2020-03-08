import player from './player';

class http extends player {
  play(board, whoami, last, callback) {
    if (whoami === 1) {
      // TODO: transpose board and last, swap players.
    }

    const body = {
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

    fetch(endpoint, request)
      .then(response => response.json())
      .then(data => {
        console.log('Expected outcome:', data.expectedOutcome);
        callback(data.move[0], data.move[1]);
      });
  }
}

export default http;
