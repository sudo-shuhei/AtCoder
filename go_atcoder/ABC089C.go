package main

import "fmt"

func main(){
  var N int
  fmt.Scan(&N)

  S := make([]string,N)
  for i:=0;i<N;i++{
    fmt.Scan(&S[i])
  }
  m:= make(map[string]int)
  for i:=0;i<N;i++{
    first_char := S[i][0:1]
    if (first_char == "M" || first_char == "A" || first_char == "R" || first_char == "C" || first_char == "H"){
      m[first_char] += 1
    }
  }
  // fmt.Println(m)
  P := [10]string{"M","M","M","M","M","M","A","A","A","R"}
  Q := [10]string{"A","A","A","R","R","C","R","R","C","C"}
  R := [10]string{"R","C","H","C","H","H","C","H","H","H"}

  result :=0
  for i:=0;i<10;i++{
    // fmt.Println(m[P[i]]*m[Q[i]]*m[R[i]])
    result += m[P[i]]*m[Q[i]]*m[R[i]]
  }
  fmt.Println(result)
}
