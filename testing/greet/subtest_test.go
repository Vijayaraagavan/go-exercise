package main

import (
	"testing"
	"fmt"
)
type testcase struct {
  arg1 int
  arg2 int
  want int
}

func TestMultiply (t *testing.T) {
  cases := []testcase{
    {2, 3, 6},
    {10, 5, 50},
    {-8, -3, 24},
    {0, 9, 1},
    {-7, 6, -42},
  }
  
  for _, tc := range cases {
    t.Run (fmt.Sprintf("%v*%v=%v", tc.arg1, tc.arg2, tc.want), func (t *testing.T) {
      got := Multiply(tc.arg1, tc.arg2)
      if got != tc.want {
        t.Fatalf("got %v; want %v\n", got, tc.want)  
      }
    })
  }
  
}
