package main
import "fmt"
import "strconv"

func sum_of_digits(i int) int{ //10進法での各桁の和を求める関数
  s := strconv.Itoa(i)
  result := 0
  for j:=0; j<= len(s)-1; j++{
    // fmt.Println(s[j:j+1])
    var k int
    k,_ = strconv.Atoi(s[j:j+1])
    result += k
  }
  return result
}

func main(){
  var(
    N int
    A int
    B int
  )
  fmt.Scan(&N)
  fmt.Scan(&A)
  fmt.Scan(&B)
  // fmt.Println(sum_of_digits(N))
  result := 0
  for i:=1; i<=N; i++{
    num := sum_of_digits(i)
    if num >= A && num<= B{
      result += i
    }
  }
  // fmt.Println(N,A,B)
  fmt.Println(result)

}
