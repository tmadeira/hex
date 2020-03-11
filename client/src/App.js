import React, { Component } from 'react';

import Board from './Board';

import http from './players/http';
import rand from './players/random';

import './App.css';

class App extends Component {
  render() {
    const n = 9;
    const delay = 10;

    const human = 'human';
    const random = new rand(n);
    const minimax = new http(n, {
      strategy: 'ab-minimax',
      depth: 5,
      heuristic: 'mindistance-bridges',
    });

    return (
      <div className="App">
        <Board delay={delay} n={n} players={[random, minimax]} />
      </div>
    );
  }
}

export default App;
