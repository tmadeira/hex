hex
==

This is an implementation of the Hex board game. According to Wikipedia,

> Hex is a strategy board game for two players played on a hexagonal grid, theoretically of any size and several possible shapes, but traditionally as an 11×11 rhombus. Players alternate placing markers or stones on unoccupied spaces in an attempt to link their opposite sides of the board in an unbroken chain. One player must win; there are no draws. The game has deep strategy, sharp tactics and a profound mathematical underpinning related to the Brouwer fixed-point theorem. It was invented in the 1940s independently by two mathematicians, Piet Hein and John Nash. The game was first marketed as a board game in Denmark under the name Con-tac-tix, and Parker Brothers marketed a version of it in 1952 called Hex; they are no longer in production. Hex can also be played with paper and pencil on hexagonally ruled graph paper.

**See it live at [hex-game.surge.sh](http://hex-game.surge.sh/).**

You can find (or create) different players (e.g. random, minimax algorithm using alpha-beta pruning) in `src/players`.

You can change the board size and the players in the game in `src/App.js`. By default, the game is 9x9 and the players are a human (player 1, red) and a minimax AI using the heuristic function in `algorithms/smallest.js` (player 2, blue).

## Development

Install [yarn](https://yarnpkg.com) and run `yarn start`.

## Deployment

Install [surge](https://surge.sh/), run `yarn build` and `surge build hex-game.surge.sh`.
