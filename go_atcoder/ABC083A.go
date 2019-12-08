package main

import (
  "fmt"
)
func main(){
  var A,B,C,D int
  fmt.Scan(&A,&B,&C,&D)
  l := A+B
  r := C+D
  if r < l{
    fmt.Println("Left")
  }else if r == l{
    fmt.Println("Balanced")
  }else{
    fmt.Println("Right")
  }
}
