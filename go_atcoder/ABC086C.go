package main //AC

import (
  "fmt"
  // "math"
)

func main(){
  var N int
  fmt.Scan(&N)
  a := make([][]int, N)
  for i:=0; i<N; i++{
    a[i] = make([]int, 3)
    for j:=0; j<3; j++{
      fmt.Scan(&a[i][j])
    }
  }
  result := 1 // okなら１　ダメなら0
  for i:= 0; i<N; i++{
    if i==0{
      t := a[i][0]
      // x := a[i][1]
      // y := a[i][2]
      if AbsDistance([]int{0,0,0}, a[i]) > t{
        result = 0
      }else if AbsDistance([]int{0,0,0}, a[i])%2 != t%2{
        result = 0
      }
      continue
    }
    t := a[i][0]-a[i-1][0]
    // fmt.Println(t)
    if AbsDistance(a[i-1],a[i]) > t{
      result = 0
    }else if AbsDistance(a[i-1], a[i])%2 != t%2{
      result = 0
    }
    // fmt.Println(AbsDistance(a[i-1], a[i]))
  }
  if result == 0{
    fmt.Println("No")
  }else if result ==1{
    fmt.Println("Yes")
  }
}

func AbsDistance(a1 []int ,a2 []int) int{ //絶対値和を求める
  x1 := a1[1]
  y1 := a1[2]
  x2 := a2[1]
  y2 := a2[2]
  x := x1-x2
  if x < 0{
    x = -x
  }
  y := y1-y2
  if y < 0{
    y = -y
  }

  return x+y

}
