package main
//答え見た
import (
  "fmt"
  "math"
)
var M int =  1000000007

func main(){
  var n int
  fmt.Scan(&n)
  a := make([]int,n)
  for i:=0;i<n;i++{
    fmt.Scan(&a[i])
  }
  m := max(a)
  half := float64(m)/2
  // fmt.Println(half)
  h := float64(m)
  for i:= 0;i<n;i++{
    tmp := math.Abs(float64(a[i]) - half)
    if tmp <= math.Abs(h - half) && a[i] != m{
      h = float64(a[i])
    }
  }
  // fmt.Println(h)
  // fmt.Println(combination(m,int(h)))
  fmt.Println(m,int(h))

}

func max(nums []int) int {
    if len(nums) == 0 {
        panic("funciton max() requires at least one argument.")
    }
    res := nums[0]
    for i := 0; i < len(nums); i++ {
        res = int(math.Max(float64(res), float64(nums[i])))
    }
    return res
}

func permutation(n int, k int) int {
    v := 1
    if 0 < k && k <= n {
        for i := 0; i < k; i++ {
            v *= (n - i)
        }
    } else if k > n {
        v = 0
    }
    return v
}

func factorial(n int) int {
    return permutation(n, n - 1)
}

func combination(n int, k int) int {
    return permutation(n, k) / factorial(k)
}
