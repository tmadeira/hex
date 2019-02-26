import React, {Component} from 'react';
import PropTypes from 'prop-types';

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

    const board = Array(props.n);
    for (let i = 0; i < props.n; i++) {
      board[i] = Array(props.n);
      for (let j = 0; j < props.n; j++) {
        board[i][j] = 0;
      }
    }

    this.state = {
      board: board,
      turn: RED,
    };
  }

  click(X, Y) {
    const board = this.state.board;

    if (board[X][Y]) {
      return;
    }

    board[X][Y] = this.state.turn;
    this.setState({
      board: board,
      turn: this.state.turn === RED ? BLUE : RED,
    });
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
        <svg version="1.1" viewBox={`-2 -2 ${w+4} ${h+4}`}>
          {
            Array.from(Array(this.props.n).keys()).map(i =>
              Array.from(Array(this.props.n).keys()).map(j =>
                <polygon className={this.color(i, j)} key={`${i} ${j}`} onClick={() => this.click(i, j)} points={this.points(i, j)} />
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
  n: PropTypes.number,
};

export default Board;