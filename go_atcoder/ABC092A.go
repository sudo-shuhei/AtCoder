package main

import "fmt"

func main(){
  var A,B,C,D int
  fmt.Scan(&A,&B,&C,&D)
  fmt.Println(min(A,B)+min(C,D))
}

func min(A int, B int)int{
  if A<=B{
    return A
  }else{
    return B
  }
}
