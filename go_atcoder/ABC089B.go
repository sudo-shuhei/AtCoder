package main

import "fmt"

func main(){
  var N int
  fmt.Scan(&N)
  S := make([]string, N)
  result := "Three"
  for i:=0;i<N;i++{
    fmt.Scan(&S[i])

  }
  for i:=0;i<N;i++{
    if S[i] == "Y"{
      result = "Four"
    }
  }
  fmt.Println(result)
}
