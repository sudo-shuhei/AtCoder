package main

import "fmt"

func main(){
  var N int
  fmt.Scan(&N)

  C := make([]int , N-1)
  S := make([]int , N-1)
  F := make([]int , N-1)
  for i :=0; i< N-1; i++{
    fmt.Scan(&C[i], &S[i], &F[i])
  }
  // fmt.Println(C,S,F)
  for i := 0; i< N; i++{ //i番目の駅について
    t := 0
    // print("i",i)
    if i == N-1{
      fmt.Println(0)
      break
    }
    for j:= i; j<N-1 ;j++{
      // print("j",j)
      if t>= S[j]{
        waitingTime := 0
        if t%F[j] == 0{
          waitingTime = 0
        }else{
          waitingTime = F[j] - t%F[j]
        }
        t += waitingTime //待ち時間
        t += C[j] //電車に乗る時間
      }else{
        t = S[j]
        t += C[j]
      }
    }
    fmt.Println(t)
  }
}
