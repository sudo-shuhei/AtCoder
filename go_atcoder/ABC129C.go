package main

import (
  "fmt"
  "bufio"
  "strconv"
  "os"
)

// var memo = make([]int, 100005)

func main() {

	// var N, M int
	// fmt.Scan(&N, &M)
	// a := make([]int, M)
  sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanWords)
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    sc.Scan()
    M, _ := strconv.Atoi(sc.Text())

  steps := make([]int, N+1)
  for i := range steps{
    steps[i] =1
  }
	for i := 0; i < M; i++ {
		sc.Scan()
    a, _ := strconv.Atoi(sc.Text())
    steps[a] = 0
	}

  for i := 2; i<=N; i++{
    steps[i] *= (steps[i-1] + steps[i-2]) %1000000007
  }
  fmt.Println(steps[N])

	// fmt.Println(a)

	// result := 1
	// for i := 0; i <= M; i++ {
	// 	if i == 0 {
	// 		result *= conb(a[0] - 1)
	// 	} else if i == M {
	// 		result *= conb(N - a[i-1] - 1)
	// 	} else {
	// 		result *= conb(a[i] - a[i-1] - 2)
	// 	}
	// 	// result *= conb(a[i]-a[i-1])
	// }
	// fmt.Println(result%1000000007)
}

// func conb(d int) int { //d段まで何通りあるか　0スタート
// 	if memo[d] != 0 {
// 		return memo[d]
// 	} else if d == 2 {
// 		memo[d] = 2
// 		return memo[d]
// 	} else if d == 1 {
// 		memo[d] = 1
// 		return memo[d]
// 	} else if d == 0 {
// 		memo[d] = 1
// 		return memo[d]
// 	} else {
// 		memo[d] = conb(d-1) + conb(d-2)
// 		return memo[d]
// 	}
// }
