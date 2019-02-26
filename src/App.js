import React, { Component } from 'react';
import './App.css';

const N = 11;
const scale = 15;

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

    this.click = this.click.bind(this);
    this.color = this.color.bind(this);

    const board = Array(N);
    for (let i = 0; i < N; i++) {
      board[i] = Array(N);
      for (let j = 0; j < N; j++) {
        board[i][j] = 0;
      }
    }

    this.state = {
      N: N,
      board: board,
      turn: RED,
    };
  }

  click(X, Y) {
    const board = this.state.board;
    board[X][Y] = this.state.turn;
    this.setState({
      board: board,
      turn: this.state.turn === RED ? BLUE : RED,
    });
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
    const w = 3 * scale * Math.sqrt(3) * this.state.N;
    const h = 3 * scale * this.state.N + (scale + 1);

    const ww = 3 * scale * Math.sqrt(3);
    const hh = 3 * scale + (scale + 1);

    return (
      <React.Fragment>
        <svg version="1.1" viewBox={`0 0 ${w} ${h}`} width={w} height={h}>
          {
            Array.from(Array(this.state.N).keys()).map(i =>
              Array.from(Array(this.state.N).keys()).map(j =>
                <polygon className={this.color(i, j)} key={`${i} ${j}`} onClick={() => this.click(i, j)} points={this.points(i, j)} />
              )
            )
          }
        </svg>
        <svg version="1.1" viewBox={`0 0 ${ww} ${hh}`} width={ww} height={hh}>
          <polygon className={this.state.turn === RED ? "red" : "blue"} points={this.points(0, 0)} />
        </svg>
      </React.Fragment>
    );
  }
}

class App extends Component {
  render() {
    return (
      <div className="App">
        <Board />
      </div>
    );
  }
}

export default App;
