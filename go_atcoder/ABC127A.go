package main

import "fmt"

func main(){
  var A,B int
  fmt.Scan(&A, &B)
  if A >= 13{
    fmt.Println(B)
  }else if A >= 6 && A <= 12{
    fmt.Println(B/2)
  }else{
    fmt.Println(0)
  }
}
