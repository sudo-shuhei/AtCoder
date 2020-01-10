package main

import "fmt"

func max(a int, b int )int{
  if a>=b{
    return a
  }else{
    return b
  }
}

func main(){
  var Q int
  fmt.Scan(&Q)
  var A,B int
  for i:=0;i<Q;i++{
    fmt.Scan(&A,&B)
    // fmt.Println(A,B)
    product := A*B
    result := 0
    for i:=0;i<max(A,B);i++{
      pow := (i+1)*(i+1)
      if pow == product{
        result += 1
      }
      if pow >= product{
        j := (product-1)/i
        if i != j{
          result += 2*i
        }else{
          result += 2*i-1
        }
        break
      }
    }
    result -= 1
    // fmt.Println(result)
    // prevj := 0
    // for i:=1;i<=product-1;i++{
    //   j := (product-1)/i
    //   // fmt.Println(i,j)
    //   if i == A || j == B{
    //     continue
    //   }else if j == prevj{
    //     continue
    //   }else{
    //     prevj = j
    //     result += 1
    //     // fmt.Println("●")
    //   }

    // }
    fmt.Println(result)
  }
}
