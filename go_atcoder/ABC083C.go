package main

import (
  "fmt"
)

func main() {
  var X, Y int
  fmt.Scan(&X,&Y)

  result := 1
  a := X
  for {
    if a*2 <= Y{
      result +=1
      a = a*2
    }else{
      break
    }
  }
    fmt.Println(result)
}
