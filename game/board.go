package game


// Player Board structure

type PlayerBoard struct {
  Score int
  Holes []int
}

func NewPlayerBoard() *PlayerBoard {
  new_board := &PlayerBoard{}
  new_board.Score = 0
  new_board.Holes = []int{ 6,6,6,6,6,6 }
  return new_board
}

