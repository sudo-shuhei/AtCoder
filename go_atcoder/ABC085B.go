// Kagami Mochi
package main

import ("fmt")

func contains(s []int, e int) bool{
  for _, v := range s{
    if e == v{
      return true
    }
  }
  return false
}

func main(){

  var N int
  fmt.Scan(&N)

  d := make([]int, N)
  for i:=0; i<N; i++{
    fmt.Scan(&d[i])
  }

  fmt.Println(kagamimochi(N, d))
}

func kagamimochi(N int, d []int) int{
  var set []int
  for i:=0; i<N; i++{
    contain := contains(set, d[i])
    if contain{
      continue
    }else{
      set = append(set, d[i])
    }
  }
  return len(set)
}
