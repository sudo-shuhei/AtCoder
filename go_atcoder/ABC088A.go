package main

import "fmt"

func main(){
  var N int
  var A int
  fmt.Scan(&N, &A)
  extra := N%500
  if extra <= A{
    fmt.Println("Yes")
  }else{
    fmt.Println("No")
  }
}
