package main

import "fmt"
const MOD = 1000000007
func main(){
  var N int
  fmt.Scan(&N)
  A := make([]int, N)
  for i :=0;i<N;i++{
    fmt.Scan(&A[i])
  }
  l := A[0]
  for i:=1;i<N;i++{
    l = lcm(l, A[i])
  }
  // fmt.Println(l)
  result := 0
  for i:=0;i<N;i++{
    result += l/A[i]
  }
  // fmt.Println(result)
  fmt.Println(result % MOD)
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func lcm(a,b int) int {
  return a*b / gcd(a,b)
}
