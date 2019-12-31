package main

import (
  "fmt"
  "strconv"
)

func main(){
  var A,B int
  fmt.Scan(&A, &B)
  result := 0
  for i:=A; i<=B; i++{
    s := strconv.Itoa(i)
    // fmt.Println(s)
    check := 1
    for j:=0; j<len(s); j++{
      // fmt.Println(j)
      if s[j] != s[len(s) - j-1]{
        check = 0
      }
    }
    if check == 1{
      result +=1
    }
  }
  fmt.Println(result)
}
