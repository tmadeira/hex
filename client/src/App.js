import React, { Component } from 'react';

import Board from './Board';

import http from './players/http';
import rand from './players/random';

import './App.css';

class App extends Component {
  render() {
    const n = 9;

     /* eslint-disable no-unused-vars */

    const human = 'human';
    const random = new rand(n);

    const mindist = d => (
      new http(n, {
        strategy: 'ab-minimax',
        depth: d,
        heuristic: 'mindistance',
      })
    );

    const bridges = d => (
      new http(n, {
        strategy: 'ab-minimax',
        depth: d,
        heuristic: 'mindistance-bridges',
      })
    );

    const both = d => (
      new http(n, {
        strategy: 'ab-minimax',
        depth: d,
        heuristic: 'mindistance-bridges-both',
      })
    );

    return (
      <div className="App">
        <Board delay={300} n={n} players={[human, both(5)]} restart={false} />
      </div>
    );
  }
}

export default App;
