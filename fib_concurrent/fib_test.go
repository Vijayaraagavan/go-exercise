package main

import (
	"testing"
  "time"
  "sync"
)
type testcase struct {  //in case we use table driven test
    data int
    check int
    wg *sync.WaitGroup
    in []int
    index int
    start time.Time
}

func TestFibrecursive (t *testing.T) {
  in := []int{10}
  data := 10
  check := data
  var wg sync.WaitGroup
  wg.Add(1)
  index := 0
  start := time.Now()

  expect := 55
  got := Fibrecursive(data, check, &wg, in, index, start)
  if expect != got {
    t.Errorf("got %v; want %v\n", got, expect)
  }
}

func BenchmarkFibrecursive (t *testing.B) {
  in := []int{40}
  data := 40
  check := data
  var wg sync.WaitGroup
  wg.Add(1)
  index := 0
  start := time.Now()

  expect := 102334155
  got := Fibrecursive(data, check, &wg, in, index, start)
  if expect != got {
    t.Errorf("got %v; want %v\n", got, expect)
  }
}

/*
  "go test" will only run test func and not benchmark
  use flags for specific tests
  
  command: go test -run=TestFibrecursive -v
      === RUN   TestFibrecursive
      --- PASS: TestFibrecursive (0.00s)
      PASS
      ok      github.com/Vijayaraagavan/rolopos.git/menuconfig/buf    0.377s
  --------------------------------------------------------------------------------
  
  command: go go test -bench=BenchmarkFibrecursive -v
      === RUN   TestFibrecursive
      --- PASS: TestFibrecursive (0.00s)
      goos: darwin
      goarch: amd64
      pkg: github.com/Vijayaraagavan/rolopos.git/menuconfig/buf
      cpu: Intel(R) Core(TM) i5-3210M CPU @ 2.50GHz
      BenchmarkFibrecursive
      BenchmarkFibrecursive-4                1        1534319665 ns/op
      PASS
      ok      github.com/Vijayaraagavan/rolopos.git/menuconfig/buf    1.702s
*/
