package main

import (
  "fmt"
)

func QuickSort(dat []int32, left int32, right int32, asc bool) {
  if right-left <= 1 {
    return
  }
  var init_pivot_i, last_i int32 = (left+right)/2, right-1
  dat[init_pivot_i], dat[last_i] = dat[last_i], dat[init_pivot_i]

  dose_swap := func(x, y int32) bool {return x > y}
  if asc {
    dose_swap = func(x, y int32) bool {return x < y}
  }
  var i int32 = left
  var j int32
  for j = left; j < last_i; j++ {
    if dose_swap(dat[j], dat[last_i]) {
      dat[i], dat[j] = dat[j], dat[i]
      i++
    }
  }
  dat[i], dat[last_i] = dat[last_i], dat[i]

  QuickSort(dat, left, i, asc)
  QuickSort(dat, i+1, right, asc)
}

func main() {
  a := []int32{3,1,2,5,7,4,6,7,9,4,8,9,5,5,3,8}
  fmt.Println(a)
  QuickSort(a, 0, int32(len(a)), true)
  fmt.Println(a)
}
