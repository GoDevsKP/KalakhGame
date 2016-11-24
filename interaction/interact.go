package interaction

import (
  "fmt"
  "time"
  "kalakh/game"
  "os"
  "os/exec"
)

func Interact(){
  curr_game := game.Game{}
  fmt.Printf("Wellcome in Kalakh\n")
  player1 := game.NewPlayer(1)
  player2 := game.NewPlayer(2)
  curr_game.Init(player1, player2)

  for {
    var input int
    Clear()
    fmt.Printf("Player's %v step.\nEnter a number [1-6] to sow from the given pit\n", curr_game.ActivePlayer.Id)

    curr_game.PrintBoard()

    fmt.Scanln(&input)
    status := curr_game.NextStep(input - 1)

    curr_game.PrintBoard()

    if status == game.Invalid {
      Clear()
      fmt.Printf("Invalid step! Try again!")
      time.Sleep(1000 * time.Millisecond)
    }
    if status == game.Win {
      fmt.Printf("Won Player#%v", curr_game.Winner.Id)
      break
    }
    if status == game.Draw {
      fmt.Printf("Draw")
      break
    }
    if status == game.TurnNextPlayer {
      curr_game.ActivePlayer = curr_game.ActivePlayer.Enemy
    }
  }
}

func Clear(){
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
}
