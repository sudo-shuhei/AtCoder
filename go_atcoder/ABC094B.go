package main
import "fmt"
func main() {
  var N,M,X int
  fmt.Scan(&N,&M,&X)
  A := make([]int,M)
  right :=0
  for i:=0;i<M;i++{
    fmt.Scan(&A[i])
    if A[i] >= X{
      right += 1
    }
  }
  left := len(A) - right
  if left < right{
    fmt.Println(left)
  }else{
    fmt.Println(right)
  }
}
