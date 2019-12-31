package main

import "fmt"

func main(){
  var c1, c2, c3 string
  fmt.Scan(&c1,&c2,&c3)
  fmt.Println(c1[0:1]+c2[1:2]+c3[2:3])
}
