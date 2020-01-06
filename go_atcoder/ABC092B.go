package main

import "fmt"

func main(){
  var N,D,X int
  fmt.Scan(&N,&D,&X)
  result := X
  A := make([]int,N)
  for i:=0;i<N;i++{
    fmt.Scan(&A[i])
    if D%A[i]==0{
      result += D/A[i]
    }else{
      result += D/A[i]+1
    }
  }
  fmt.Println(result)
}
