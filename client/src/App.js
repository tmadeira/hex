import React, { Component } from 'react';

import Board from './Board';

import http from './players/http';
import rand from './players/random';

import './App.css';

class App extends Component {
  render() {
    const n = 9;
    const delay = 300;

    const human = 'human';
    const random = new rand(n);

    const mindist = new http(n, {
      strategy: 'ab-minimax',
      depth: 5,
      heuristic: 'mindistance',
    });
    const bridges = new http(n, {
      strategy: 'ab-minimax',
      depth: 5,
      heuristic: 'mindistance-bridges',
    });
    const both = new http(n, {
      strategy: 'ab-minimax',
      depth: 5,
      heuristic: 'mindistance-bridges-both',
    });

    return (
      <div className="App">
        <Board delay={delay} n={n} players={[random, both]} />
      </div>
    );
  }
}

export default App;
