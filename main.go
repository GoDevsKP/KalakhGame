package main

import (
  "kalakh/board"
)

func main() {
  player1 := board.NewPlayer(1)
  player2 := board.NewPlayer(2)

  game := board.Game{}
  game.Init(player1, player2)
  stepState := game.NextStep(1)
  game.UpdateState(stepState)
  game.NextStep(2);
  // player2.MakeStep(2)
}
