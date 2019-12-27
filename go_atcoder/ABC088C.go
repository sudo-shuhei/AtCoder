package main

import "fmt"

func main(){
  // line := make([]int, 3)
  // arr := make([][]int, 3)
  // arr[0] = line
  // arr[1] = line
  // arr[2] = line
  arr := make([][]int, 3)
  for i:=0; i<3; i++{
    arr[i] = make([]int, 3)
  }
  // fmt.Println(arr)
  fmt.Scan(&arr[0][0], &arr[0][1], &arr[0][2])
  fmt.Scan(&arr[1][0], &arr[1][1], &arr[1][2])
  fmt.Scan(&arr[2][0], &arr[2][1], &arr[2][2])
  // fmt.Println(arr)
  line1 := make([]int, 3)
  line2 := make([]int, 3)
  line3 := make([]int, 3)
  line4 := make([]int, 3)

  for i:=0; i<3; i++{
    line1[i] = arr[0][i] - arr[1][i]
    line2[i] = arr[1][i] - arr[2][i]
  }
  for i:=0; i<3; i++{
    line3[i] = arr[i][0] - arr[i][1]
    line4[i] = arr[i][1] - arr[i][2]
  }
  result := 1
  val1 := line1[0]
  val2 := line2[0]
  val3 := line3[0]
  val4 := line4[0]

  for i:=0; i<3; i++{
    if line1[i] != val1{
      result = 0
    }
    if line2[i] != val2{
      result = 0
    }
    if line3[i] != val3{
      result = 0
    }
    if line4[i] != val4{
      result = 0
    }
  }
  if result == 0{
    fmt.Println("No")
  }else{
    fmt.Println("Yes")
  }

}
