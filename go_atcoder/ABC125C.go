package main

import (
  "fmt"
)

// const max_int int =1000000000
var memo [][]int
// var i int = 0

func gcd(a, b int) int{
    if b == 0 {
        return a
    }
    if a < b{
      if memo[a][b] != 0{
        return memo[a][b]
      }
      memo[a][b] = gcd(b,a%b)
      return memo[a][b]
    }else{
      if memo[b][a] != 0{
        return memo[a][b]
      }
      memo[b][a] = gcd(b,a%b)
      return memo[b][a]
    }
    // return memo[a][b]
}

func main(){
  var N int
  fmt.Scan(&N)

  A := make([]int, N)
  for i:=0; i<N; i++{
    fmt.Scan(&A[i])
  }
  // fmt.Println(gcd(20,12))
  max_gcd := 0
  for i:=0; i<N; i++{
    //A[i]をスキップしてGCDを求める
    // fmt.Println(i)
    gcd_tmp := 0
    for j:=0; j<N; j++{
      if j == i{
        continue
      }
      if gcd_tmp == 0{
        gcd_tmp = A[j]
      }
      gcd_tmp = int(gcd(gcd_tmp,A[j]))
    }
    if gcd_tmp > max_gcd {
      max_gcd = gcd_tmp
    }
  }
  fmt.Println(max_gcd)
}
