Switch in go are actual switch

unlike other programming languages where without break statement it runs the rest of the code.
Go runs only the match case.

switch os := runtime.GOOS; os {
  case "darwin":
    // do something
  case "linux":
    // do something else
  default:
    // do nothing
}

evaulation order is top to bottom.


switch without condition

switch {
  // ...
}

is used to write long clean if else

switch { // true
  case isWin():
    // do something
  case isLose():
    // do something
  default:
    // anything
}