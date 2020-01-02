package main

import "fmt"

func main(){
  var N int
  a := make([]int,N)
  b := make([]int,N)
  c := make([]int,N)
  d := make([]int,N)
  for i:=0;i<N;i++{
    fmt.Scan(&a[i])
    fmt.Scan(&b[i])
  }
  for i:=0;i<N;i++{
    fmt.Scan(&c[i])
    fmt.Scan(&d[i])
  }
  
}
