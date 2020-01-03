package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "strconv"
)
var n, k int
var max_working_people int
// var zeros []int

func main() {
	lines := getStdin()
	// for i, v := range lines {
	// 	fmt.Printf("line[%d]: %s\n", i, v)
  // fmt.Println(lines[0],lines[1])
  nk := strings.Fields(lines[0])
  n,_ = strconv.Atoi(nk[0])
  k,_ = strconv.Atoi(nk[1])
  s := lines[1]
  l := make([]int, len(s))
  for i :=0; i< len(s); i++{
    if s[i:i+1] == "S"{
      l[i] = 1
    }else{
      l[i] = 0
    }
  }
  // fmt.Println(l)
  max_working_people = 0
  // zeros := []int{0,0,0}
  rec(l,k)
  fmt.Println(max_working_people)

	}

func rec(l []int, r int){
  // fmt.Println(l,r)
  if r ==0{
    s := sum(l)
    working_people := n - s
    if working_people > max_working_people{
      max_working_people = working_people
    }
    return
  }
  r -=1
  index := search(l)
  // fmt.Println(index)

  for i:=0;i<len(index);i++{
    // fmt.Println(index[i])
    l2 := make([]int,len(l))
    copy(l2,l)
    // fmt.Println("l2",l2)
    l2[index[i]] = 0
    l2[index[i]+1] = 0
    l2[index[i]+2] = 0
    rec(l2,r)
  }
}

func search(l []int) []int{
  d := make(map[int][]int)
  d[3] = make([]int,0)
  d[2] = make([]int,0)
  d[1] = make([]int,0)
  d[0] = make([]int,0)

  for i:=0; i<len(l)-2; i++{
    slice := l[i:i+3]
    sum_slice := sum(slice)
    if sum_slice == 3{
      d[3] = append(d[3], i)
    }else if sum_slice == 2 && slice[1] != 0{
      d[2] = append(d[2],i)
    }else if sum_slice == 1 && slice[1] != 0{
      d[1] = append(d[1],i)
    }else{
      d[0] = append(d[0],i)
    }
  }
  // fmt.Println(d)
  if len(d[3]) != 0{
    return d[3]
  }else if len(d[2]) != 0{
    return d[2]
  }else if len(d[1]) != 0{
    return d[1]
  }else{
    return d[0]
  }
}

func sum(l []int) int {
  result := 0
  for i:=0;i<len(l);i++{
    result += l[i]
  }
  return result
}

func getStdin() []string {
	stdin := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		lines = append(lines, stdin.Text())
	}
	return lines
}
