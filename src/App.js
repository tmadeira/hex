import React, { Component } from 'react';

import Board from './Board';

import random from './players/random';

import './App.css';

class App extends Component {
  render() {
    const n = 9;

    const a = new random(n);
    const b = 'human';

    return (
      <div className="App">
        <Board n={n} players={[a, b]} />
      </div>
    );
  }
}

export default App;
