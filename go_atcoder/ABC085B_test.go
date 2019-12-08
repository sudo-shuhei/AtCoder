package main

import (
  "testing"
)

func TestKagamimochi(t *testing.T){
  N := 4
  d := []int{10,8,6,6}
  result := kagamimochi(N,d)
  expected := 3
  if result != expected{
    t.Errorf("expected %d but got %d", expected, result)
  }
}

func ExampleAdd() {
    N:= 5
    d:= []int{2, 4, 20, 20, 5}
    result := kagamimochi(N,d)
    fmt.Println(result)
    // Output: 5
}
