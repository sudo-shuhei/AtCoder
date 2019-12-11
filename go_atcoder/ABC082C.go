package main

import (
  "fmt"
)

func main() {
  var N int
  fmt.Scan(&N)
  a := make([]int, N)
  for i:= range(a){
    fmt.Scan(&a[i])
  }
  // fmt.Println(a)
  m := map[int]int{}
  for i:= range(a){
    m[a[i]] +=1
  }
  // fmt.Println(m)

  result := 0
  for k, v := range(m){
    if k < v{
      result += v-k
    }else if k == v{
      continue
    }else{
      result += v
    }
  }
  fmt.Println(result)
}
