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

    const both1 = new http(n, {
      strategy: 'ab-minimax',
      depth: 1,
      heuristic: 'mindistance-bridges-both',
    });

    const both = new http(n, {
      strategy: 'ab-minimax',
      depth: 5,
      heuristic: 'mindistance-bridges-both',
    });

    return (
      <div className="App">
        <Board delay={300} n={n} players={[human, both]} restart={false} />
      </div>
    );
  }
}

export default App;
