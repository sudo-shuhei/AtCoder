package main

import (
  "fmt"
  "sort"
  )

func main(){
  var a,b,c int
  fmt.Scan(&a,&b,&c)
  s := []int{a,b,c}
  // fmt.Println(s)
  sort.Ints(s)
  // fmt.Println(s)
  max := s[2]
  mid := s[1]
  min := s[0]

  result := 0
  if mid%2 == min%2{
    result += (mid-min)/2
    result += max - mid
  }else{
    max += 1
    min += 1
    result +=1
    result += (mid-min)/2
    result += max - mid
  }
  fmt.Println(result)
}
