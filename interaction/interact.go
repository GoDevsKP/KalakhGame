package interaction

import (
  "fmt"
  "time"
  "kalakh/board"
  "os"
  "os/exec"
)

func Interact(){
  game := board.Game{}
  fmt.Printf("Wellcome in Kalakh\n")
  player1 := board.NewPlayer(1)
  player2 := board.NewPlayer(2)
  game.Init(player1, player2)

  for {
    var input int
    Clear()
    fmt.Printf("Player's %v step.\nEnter a number [1-6] to sow from the given pit\n", game.ActivePlayer.Id)

    game.PrintBoard()

    fmt.Scanln(&input)
    status := game.NextStep(input - 1)

    game.PrintBoard()

    if status == board.Invalid {
      Clear()
      fmt.Printf("Invalid step! Try again!")
      time.Sleep(1000 * time.Millisecond)
    }
    if status == board.Win {
      fmt.Printf("Won Player#%v", game.Winner.Id)
      break
    }
    if status == board.Draw {
      fmt.Printf("Draw")
      break
    }
    if status == board.TurnNextPlayer {
      game.ActivePlayer = game.ActivePlayer.Enemy
    }
  }
}

func Clear(){
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
}
