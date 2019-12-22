package main

import (
  "fmt"
)

func main() {
  var X,A,B int
  fmt.Scan(&X,&A,&B)
  fmt.Println((X-A)%B)
}
