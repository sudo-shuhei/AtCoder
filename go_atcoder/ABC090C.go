package main

import "fmt"

func main(){
  var N,M int
  fmt.Scan(&N, &M)
  if N==1 && M==1{
    fmt.Println(1)
  }else if N == 1{
    fmt.Println(M-2)
  }else if M ==1{
    fmt.Println(N-2)
  }else{
    fmt.Println((M-2)*(N-2))
  }
}
