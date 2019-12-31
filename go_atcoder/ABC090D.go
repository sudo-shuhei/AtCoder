package main

import "fmt"

func main(){
  var N,K int
  fmt.Scan(&N,&K)
  //K以上のbに対して、あまりがK以上になるaがいくつあるか
  //c := N/b cb <= N < (c+1)b
  //c(b-K) + (N-cb)-K if N-cb >= K else 0
  result := 0
  for b := K+1; b<=N; b++{
    c := N/b
    d := 0

    if (N - c*b) >= K && N- c*b > 0{
      d = (N- c*b) - K +1
    }else{
      d = 0
    }
    fmt.Println(b,c,d, c*(b-K) + d)
    result += c*(b-K) + d
  }
  fmt.Println(result)
}
