package game

import (
  "fmt"
)

// Game

type Game struct {
  ActivePlayer *Player
  Status GameStatus
  Winner *Player
}


func (g *Game) Init(p1, p2 *Player) {
  p1.Board = NewPlayerBoard()
  p2.Board = NewPlayerBoard()
  p1.Enemy, p2.Enemy = p2, p1
  g.ActivePlayer = p1
}


func (g *Game) NextStep(from int) StepStatus {
  actor := g.ActivePlayer

  if !validate_step(from, actor.Board) {
    return Invalid
  }

  num_stones := actor.Board.Holes[from]
  actor.Board.Holes[from] = 0
  pos := from + 1

  for ; num_stones > 0; {
    // Firstly, fill actor board with stoles
    pos, num_stones = fill_stones(actor.Board, pos, num_stones)

    // If actor finished his move at own empty pit,
    // sow the pits from enimie's opposite pit
    if num_stones == 0 && actor.Board.Holes[pos] == 1 {
      actor.Board.Score += actor.Enemy.Board.Holes[USER_BOARD_LENGTH - pos - 1] + 1
      actor.Enemy.Board.Holes[USER_BOARD_LENGTH - pos - 1] = 0
      actor.Board.Holes[pos] = 0
    }

    // If stones left, fill actor kalakh
    if num_stones > 0 {
      actor.Board.Score++
      if CheckForWin(actor.Board.Score) {
        g.Winner = actor
        return Win
      }

      num_stones--
      pos = 0

      if num_stones == 0 {
        return Repeat
      }
    }

    if num_stones > 0 {
      _, num_stones = fill_stones(actor.Enemy.Board, pos, num_stones)
      pos = 0
    }


    // check for win, if one of players has no pits
    if get_total_board_score(actor.Board) == 0 || get_total_board_score(actor.Enemy.Board) == 0 {
      actor.Enemy.Board.Score += get_total_board_score(actor.Enemy.Board)
      actor.Board.Score += get_total_board_score(actor.Board)
      fill_board_with_zeros(actor.Enemy.Board)
      fill_board_with_zeros(actor.Board)

      if actor.Board.Score > actor.Enemy.Board.Score {
        g.Winner = actor
      }
      if actor.Board.Score < actor.Enemy.Board.Score {
        g.Winner = actor.Enemy
      }
      if actor.Board.Score == actor.Enemy.Board.Score {
        return Draw
      }

      return Win
    }
  }

  return TurnNextPlayer
}


func fill_stones(board *PlayerBoard, from int, stones int) (int,int){
  var left = stones
  var pos int
  for pos = from; pos < USER_BOARD_LENGTH && left > 0; pos++ {
    board.Holes[pos] += 1
    left--
  }
  return pos - 1, left
}


func get_total_board_score(board *PlayerBoard) int {
  var sum int = 0
  for  i := 0; i < USER_BOARD_LENGTH; i++ {
    sum += board.Holes[i]
  }
  return sum
}


func fill_board_with_zeros(board *PlayerBoard) {
  for i := 0; i < USER_BOARD_LENGTH; i++ {
    board.Holes[i] = 0
  }
}


func validate_step(from int, board *PlayerBoard) bool {
  if from > USER_BOARD_LENGTH - 1 || from < 0 {
    return false
  }
  if board.Holes[from] == 0 {
    return false
  }
  return true
}


func (g *Game) PrintBoard(){

  fmt.Printf("%v\n", reverseInts(g.ActivePlayer.Enemy.Board.Holes))
  fmt.Printf("%v\n", g.ActivePlayer.Board.Holes)
  fmt.Printf("%v-", g.ActivePlayer.Board.Score)
  fmt.Printf("%v\n***\n", g.ActivePlayer.Enemy.Board.Score)
}

func reverseInts(input []int) []int {
  if len(input) == 0 {
    return input
  }
  return append(reverseInts(input[1:]), input[0])
}
