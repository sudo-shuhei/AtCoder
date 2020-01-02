package main

import "fmt"

func main(){
  var N,M int
  fmt.Scan(&N)
  s := make([]string, N)
  m := make(map[string]int)
  for i:=0;i<N;i++{
    fmt.Scan(&s[i])
    m[s[i]] +=1
  }
  fmt.Scan(&M)
  t := make([]string,M)
  for i:=0;i<M;i++{
    fmt.Scan(&t[i])
    m[t[i]] -=1
  }
  result := 0
  for _,v:= range m{
    if v > result{
      result = v
    }
  }
  fmt.Println(result)
}
