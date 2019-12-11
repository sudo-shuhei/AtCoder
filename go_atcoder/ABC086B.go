package main

import (
  "fmt"
  "strconv"
  "math"
)

func main() {
  var a, b int
  fmt.Scan(&a, &b)
  num, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
  // fmt.Println(num)
  root := math.Sqrt(float64(num))
  // fmt.Println(root)
  if math.Ceil(root) == root{
    fmt.Println("Yes")
  }else{
    fmt.Println("No")
  }
}
