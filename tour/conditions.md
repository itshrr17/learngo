Conditional

like for, if statement can start with short statement

short statement is v := 10, to execute before the condition
variables declared inside the short statement are only valid only inside the if block

if v := Math.Pow(x, n); v < lim {
  return v
} 
// example of else if
else if v < lim2 {
  return v
} 

else {
  return lim

}