package board

import (
  "fmt"
)

type StepStatus int

const (
  TurnNextPlayer StepStatus = 1
  Repeat StepStatus = 2
  Win StepStatus = 3
)


type GameStatus int

const (
  PLAYING GameStatus = 1
  END GameStatus = 2
)


type Game struct {
  ActivePlayer *Player
  Status GameStatus
}

func (g *Game) Init(p1, p2 *Player) {
  p1.Board = NewPlayerBoard()
  p2.Board = NewPlayerBoard()
  p1.Enemy, p2.Enemy = p2, p1
  g.ActivePlayer = p1
}


func (g *Game) NextStep(from int) StepStatus {
  actor := g.ActivePlayer
  num_stones := actor.Board.holes[from]
  actor.Board.holes[from] = 0
  pos := from + 1

  for ; num_stones > 0; {
    pos, num_stones = fill_stones(actor.Board, pos, num_stones)

    if num_stones == 0 && actor.Board.holes[pos] == 1 {
      fmt.Printf("pos =%v", pos)
      actor.Board.score += actor.Enemy.Board.holes[6-pos-1] + 1
      actor.Enemy.Board.holes[6-pos-1] = 0
      actor.Board.holes[pos] = 0
    }
    if num_stones > 0 {
      actor.Board.score++
      if actor.Board.score > 36 {
        return Win
      }
      num_stones--
      pos = 0

      if num_stones == 0 {
        fmt.Printf("Repeat")
        return Repeat
      }
    }
    if num_stones > 0 {
      pos, num_stones = fill_stones(actor.Enemy.Board, pos, num_stones)
      pos = 0
    }

  }


  fmt.Printf("%v\n", actor.Board.holes)
  fmt.Printf("%v\n", actor.Enemy.Board.holes)
  fmt.Printf("%v-", actor.Board.score)
  fmt.Printf("%v\n***\n", actor.Enemy.Board.score)

  return TurnNextPlayer

}


func fill_stones(board *PlayerBoard, from int, stones int) (int,int){
  var left = stones
  var pos int
  for pos = from; pos < 6 && left > 0; pos++ {
    board.holes[pos] += 1
    left--
  }
  return pos - 1, left

}

func (g *Game) UpdateState(status StepStatus){
  if status == Win {
    fmt.Printf("Win")
  }
  if status == TurnNextPlayer {
    g.ActivePlayer = g.ActivePlayer.Enemy

  }
}



// Player Board structure

type PlayerBoard struct {
  score int
  holes [6]int
}

func NewPlayerBoard() *PlayerBoard {
  new_board := &PlayerBoard{}
  new_board.score = 0
  new_board.holes = [6]int{ 6,13,0,6,6,6 }
  return new_board
}


// Player structure

type Player struct {
  id int
  Enemy *Player
  Board *PlayerBoard
}


func NewPlayer(id int) *Player {
  new_player := &Player{ id: id }
  return new_player
}
