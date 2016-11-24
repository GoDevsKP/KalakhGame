package game


// Step state

type StepStatus int

const (
  TurnNextPlayer StepStatus = 1
  Repeat StepStatus = 2
  Win StepStatus = 3
  Draw StepStatus = 4
  Invalid StepStatus = 5
)
