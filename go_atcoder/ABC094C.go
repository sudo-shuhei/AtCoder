package main

import (
  "fmt"
  "sort"
)

func main(){
  var N int
  fmt.Scan(&N)
  X := make([]int, N)
  for i:=0;i<N;i++{
    fmt.Scan(&X[i])
  }
  Y := make([]int,N)
  copy(Y,X)
  sort.Ints(X)
  // fmt.Println(X)
  // fmt.Println(Y)
  low := X[(N-1)/2]
  high := X[N/2]
  // fmt.Println(low,high)
  for i:=0;i<N;i++{
    if Y[i] <= low{
      fmt.Println(high)
    }else{
      fmt.Println(low)
    }
  }
}
