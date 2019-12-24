package main

import "fmt"

func main () {
  var A,B int
  var S string
  fmt.Scan(&A,&B,&S)

  // if S[A+1:A+2] == "-"{
  //   for
  // }
  result := 1
  for i:=0; i<A+B+1; i++ {
    if i == A{
      if S[i:i+1] != "-"{
        result = 0
      }
    } else {
      if S[i:i+1] == "-"{
        result = 0
      }
    }
  }
  if result == 0{
    fmt.Println("No")
  }else{
    fmt.Println("Yes")
  }
  
}
