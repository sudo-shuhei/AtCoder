package main //ソート例

import "fmt"
import "sort"

func main(){
  var N int
  fmt.Scan(&N)

  a := make([]int, N)
  for i:=0; i<N; i++{
    fmt.Scan(&a[i])
  }

  // sort.Slice(a, func(i,j int) bool{ return a[i] < a[j]})
  sort.Sort(sort.Reverse(sort.IntSlice(a)))
  // fmt.Println(N, a)
  i:=0
  alice :=0
  bob :=0
  for i<=N{
    if i%2 == 0{
      alice += a[i]
    }else{
      bob += a[i]
    }
    i++
  }
  fmt.Println(alice-bob)
}
