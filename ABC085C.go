package main

import (
  "fmt"
)

func main(){
  var N int
  var Y int
  fmt.Scan(&N, &Y)

  for i:=0; i<=N; i++{
    // 10000円札がi枚
    for j:=0; j<= N-i; j++{
      // 5000円札がj枚
      k:= N-i-j // 1000円札k枚
      if 10000*i+5000*j+1000*k == Y{
        fmt.Printf("%d %d %d",i,j,k)
        return
      }
    }
  }
  fmt.Println("-1 -1 -1")
}
