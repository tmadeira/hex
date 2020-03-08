import React, { Component } from 'react';

import Board from './Board';

import http from './players/http';

import './App.css';

class App extends Component {
  render() {
    const n = 5;
    const delay = 10;

    const a = 'human';
    const b = new http(n);

    return (
      <div className="App">
        <Board delay={delay} n={n} players={[a, b]} />
      </div>
    );
  }
}

export default App;
