package game

func CheckForWin(actorScore int) bool {
  if actorScore > TOTAL_POINTS_NUM / 2 {
    return true
  } else {
    return false
  }
}


