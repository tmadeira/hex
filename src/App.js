import React, { Component } from 'react';

import Board from './Board';

import minimax from './players/minimax';
import {heuristic} from './algorithms/smallest';

import './App.css';

class App extends Component {
  render() {
    const n = 9;
    const delay = 10;

    const a = 'human';
    const b = new minimax(n, heuristic);

    return (
      <div className="App">
        <Board delay={delay} n={n} players={[a, b]} />
      </div>
    );
  }
}

export default App;
