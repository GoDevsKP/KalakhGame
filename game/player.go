package game


// Player structure

type Player struct {
  Id int
  Enemy *Player
  Board *PlayerBoard
}


func NewPlayer(id int) *Player {
  new_player := &Player{ Id: id }
  return new_player
}
