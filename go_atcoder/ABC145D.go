package main

import (
  "fmt"
  "math"
)



func main(){
  var X,Y int
  fmt.Scan(&X,&Y)
  // maxnum := max(X,Y)
  // n + 2m = X
  // 2n + m = Y
  // 4n + 2m = 2Y
  // 3n = 2Y -X
  // n = (2Y-X)/3
  // m = X - n /2 = X -(2Y-X)/6
  n := (2*Y - X)/3
  m := Y - 2*n
  fmt.Println(n,m)
  big := int(math.Pow(10,9)+7)
  if n <0 || m < 0{
    fmt.Println(0)
  }else{
  fmt.Println((factorial(n+m)/(factorial(n)*factorial(m)))%big)
  }
}

func factorial(n int) int{
  if n == 0{
    return 1
  }else if n ==1{
    return 1
  }
  return n * factorial(n-1)
}

func max(x int, y int) int{
  if x>= y{
    return x
  }else{
    return y
  }
}
