package main

import "fmt"

func main(){

  var N int
  fmt.Scan(&N)

  ii := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&ii[i])
    }

  result := 10000000000
  for _,i := range ii {
    // fmt.Printf("%d\n",i)
    power:= 0
    for i%2 ==0 {
      // fmt.Println(i)
      i /= 2
      power +=1
    }
    if power <= result{
      result = power
    }
  }
  fmt.Println(result)
}
