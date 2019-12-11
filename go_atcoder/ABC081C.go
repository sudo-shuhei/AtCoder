package main

import (
  "fmt"
  "sort"
)

func main(){
  var N, K int
  fmt.Scan(&N, &K)
  A := make([]int, N)
  for i:=0; i<N; i++{
    fmt.Scan(&A[i])
  }
  // fmt.Println(A)
  sort.Ints(A)
  // fmt.Println(A)

  // m = map[int]int{}
  m := make(map[int]int, N)
  for i:=0; i<N; i++{
    m[A[i]] += 1
  }
  // fmt.Println(m)

  l := []int{}
  for _, value := range m {
    l = append(l, value)
  }
  // fmt.Println(l)
  sort.Ints(l)
  // fmt.Println(l)
  kinds := len(l)
  ans := 0
  for i:=0; i<(kinds-K); i++{
    ans += l[i]
  }
  fmt.Println(ans)
}
