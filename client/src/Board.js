import React, {Component} from 'react';
import PropTypes from 'prop-types';

import winner from './algorithms/winner';

const scale = 30;

const base = [
  {x: scale * Math.sqrt(3), y: 4 * scale},
  {x: 0, y: 3 * scale},
  {x: 0, y: scale},
  {x: scale * Math.sqrt(3), y: 0},
  {x: 2 * scale * Math.sqrt(3), y: scale},
  {x: 2 * scale * Math.sqrt(3), y: 3 * scale},
];

const RED = 1;
const BLUE = 2;

class Board extends Component {
  constructor(props) {
    super(props);

    this.boundary = this.boundary.bind(this);
    this.click = this.click.bind(this);
    this.color = this.color.bind(this);
    this.maybePlay = this.maybePlay.bind(this);

    const board = Array(props.n + 1);
    for (let i = 0; i < props.n + 1; i++) {
      board[i] = Array(props.n + 1);
      for (let j = 0; j < props.n + 1; j++) {
        if (i === props.n) {
          board[i][j] = BLUE;
        } else if (j === props.n) {
          board[i][j] = RED;
        } else {
          board[i][j] = 0;
        }
      }
    }

    this.state = {
      blocked: true,
      board: board,
      last: [-1, -1],
      turn: RED,
    };

    this.baseState = JSON.parse(JSON.stringify(this.state));
  }

  componentDidMount() {
    this.maybePlay(this.state.turn, null);
  }

  maybePlay(player, last) {
    const other = player === RED ? BLUE : RED;
    if (typeof this.props.players[player - 1] !== 'string') {
      window.setTimeout(() => {
        this.props.players[player - 1].play(this.state.board, player, last, (x, y) => {
          this.setState({blocked: false});
          this.click(x, y);
        });
      }, typeof this.props.players[other - 1] !== 'string' ? this.props.delay : 0);
    } else {
      this.setState({blocked: false});
    }
  }

  click(X, Y) {
    const board = this.state.board;

    if (this.state.blocked || board[X][Y]) {
      return;
    }

    board[X][Y] = this.state.turn;

    const turn = this.state.turn === RED ? BLUE : RED;
    this.setState({blocked: true, board, turn, last: [X, Y]});

    window.setTimeout(() => {
      if (winner(this.state.board)) {
        if (this.props.restart) {
          window.setTimeout(() => {
            this.setState(JSON.parse(JSON.stringify(this.baseState)));
            this.componentDidMount();
          }, 1000);
        } else {
          alert('Game over!');
        }
        return;
      }

      this.maybePlay(turn, [X, Y]);
    }, 100);
  }

  boundary(which) {
    let begin = {}, dx = [], dy = [];
    switch (which) {
      case 0:
      default:
        begin = {x: 0, y: scale};
        dx = [Math.sqrt(3) * scale, Math.sqrt(3) * scale];
        dy = [-scale, scale];
        break;
      case 1:
        begin = {x: scale * Math.sqrt(3) * (this.props.n - 1), y: scale + 3 * scale * this.props.n - scale};
        dx = [Math.sqrt(3) * scale, Math.sqrt(3) * scale];
        dy = [scale, -scale];
        break;
      case 2:
        begin = {x: 0, y: scale};
        dx = [0, Math.sqrt(3) * scale];
        dy = [2 * scale, scale];
        break;
      case 3:
        begin = {x: 2 * this.props.n * Math.sqrt(3) * scale, y: scale};
        dx = [0, Math.sqrt(3) * scale];
        dy = [2 * scale, scale];
        break;
    }
    let points = '';
    let x = begin.x, y = begin.y;
    const count = which < 2 ? 2 * this.props.n + 1 : 2 * this.props.n;
    for (let i = 0; i < count; i++) {
      if (i !== 0) {
        points += ', ';
      }
      points += `${x} ${y}`;
      x += dx[i % 2];
      y += dy[i % 2];
    }
    return points;
  }

  points(X, Y) {
    let str = "";
    for (let i = 0; i < base.length; i++) {
      const x = base[i].x + 2 * scale * Math.sqrt(3) * X + scale * Math.sqrt(3) * Y;
      const y = base[i].y + 3 * scale * Y;
      if (i !== 0) {
        str += ", ";
      }
      str += `${x} ${y}`;
    }
    return str;
  }

  color(i, j) {
    if (this.state.board[i][j] === RED) {
      return "red";
    }

    if (this.state.board[i][j] === BLUE) {
      return "blue";
    }

    return "white";
  }

  render() {
    const w = 3 * scale * Math.sqrt(3) * this.props.n - Math.sqrt(3) * scale;
    const h = 3 * scale * this.props.n + (scale + 1);

    return (
      <React.Fragment>
        <svg version="1.1" viewBox={`-2 -2 ${w+4} ${h+4}`} width={w+6} height={h+6}>
          {
            Array.from(Array(this.props.n).keys()).map(i =>
              Array.from(Array(this.props.n).keys()).map(j =>
                <polygon
                  className={this.color(i, j) + (i === this.state.last[0] && j === this.state.last[1] ? ' last' : '')}
                  key={`${i} ${j}`}
                  data-coords={`${i} ${j}`}
                  onClick={() => this.click(i, j)}
                  points={this.points(i, j)}
                />
              )
            )
          }
          <polyline style={{strokeWidth: '4px'}} points={this.boundary(0)} fill="none" stroke="red" />
          <polyline style={{strokeWidth: '4px'}} points={this.boundary(1)} fill="none" stroke="red" />
          <polyline style={{strokeWidth: '4px'}} points={this.boundary(2)} fill="none" stroke="blue" />
          <polyline style={{strokeWidth: '4px'}} points={this.boundary(3)} fill="none" stroke="blue" />
        </svg>
      </React.Fragment>
    );
  }
}

Board.propTypes = {
  delay: PropTypes.number,
  n: PropTypes.number,
  players: PropTypes.array,
  restart: PropTypes.bool,
};

export default Board;
