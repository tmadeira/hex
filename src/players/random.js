import player from './player';

class random extends player {
  play(board, _whoami, _last, callback) {
    const chosen = this.randomMove(board);
    callback(chosen[0], chosen[1]);
  }
}

export default random;
