package main

import "fmt"

func main() {
  var N int
  fmt.Scan(&N)
  line1 := make([]int, N)
  line2 := make([]int, N)

  for i:= 0; i<N; i++{
    fmt.Scan(&line1[i])
  }
  for i:=0; i<N; i++{
    fmt.Scan(&line2[i])
  }
  // fmt.Println(line1)
  // fmt.Println(line2)
  result := 0
  sum := 0
  for i:=0; i<N; i++{ //i番目で下に降りる
    sum = 0
    for j:=0; j<=i; j++{
      sum += line1[j]
    }
    for k:=i; k<N; k++{
      sum += line2[k]
    }
    // fmt.Println(sum)
    if sum >= result{
        result = sum
    }
  }
  fmt.Println(result)
}
