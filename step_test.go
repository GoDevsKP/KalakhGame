package main

import (
	"testing"
	"kalakh/game"
)

func TestStepInKalakh(t *testing.T) {
	test_game := game.Game{}
	player1 := game.NewPlayer(1)
	player2 := game.NewPlayer(2)
	test_game.Init(player1, player2)
	a := test_game.ActivePlayer
	status := test_game.NextStep(0)
	if test_game.ActivePlayer != a{
		t.Errorf("After being hit in kalah expected player %v and status %v but got player %v", a.Id,status,test_game.ActivePlayer.Id)
	}


}

func TestCapture(t *testing.T){
	test_game := game.Game{}
	player1 := game.NewPlayer(1)
	player2 := game.NewPlayer(2)
	test_game.Init(player1, player2)
	test_game.ActivePlayer.Board.Holes[0]=0
	test_game.ActivePlayer.Board.Holes[1]=12
	test_game.NextStep(1)
	if test_game.ActivePlayer.Board.Score!= 9{
		t.Errorf("Capture did not happend.Score must be 9 but it is  %v", test_game.ActivePlayer.Board.Score)
	}
}

func TestWin(t *testing.T){
	test_game := game.Game{}
	player1 := game.NewPlayer(1)
	player2 := game.NewPlayer(2)
	test_game.Init(player1, player2)
	test_game.ActivePlayer.Board.Score=36
	status := test_game.NextStep(3)
	if status != game.Win{
		t.Errorf("Game did not end. Expected status 3 but got %v", status)
	}
}

func TestStepValidation(t *testing.T){
  test_game := game.Game{}
  player1 := game.NewPlayer(1)
  player2 := game.NewPlayer(2)
  test_game.Init(player1, player2)
  test_game.ActivePlayer.Board.Holes[1] = 0
  status := test_game.NextStep(1)
  if status != game.Invalid {
    t.Errorf("Invalid step, expected Invalid step status, but got %v", status)
  }
}
