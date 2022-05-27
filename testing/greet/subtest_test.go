//Table driven test using subtest
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
    t.Run (fmt.Sprintf("%v*%v=%v", tc.arg1, tc.arg2, tc.want), func (t *testing.T) {  	//first argument takes name of subtest
      got := Multiply(tc.arg1, tc.arg2)
      if got != tc.want {
        t.Fatalf("got %v; want %v\n", got, tc.want)  
      }
    })
  }
}

/*
output:
command : "go test"
	PASS
	ok      github.com/Vijayaraagavan/rolopos.git/menuconfig/buf    0.371s
	-------------------------------------------------------------------------
	
command : "go test -v"
	=== RUN   TestMultiply
	=== RUN   TestMultiply/2*3=6
	=== RUN   TestMultiply/10*5=50
	=== RUN   TestMultiply/-8*-3=24
	=== RUN   TestMultiply/0*9=0
	=== RUN   TestMultiply/-7*6=-42
	--- PASS: TestMultiply (0.00s)
	    --- PASS: TestMultiply/2*3=6 (0.00s)
	    --- PASS: TestMultiply/10*5=50 (0.00s)
	    --- PASS: TestMultiply/-8*-3=24 (0.00s)
	    --- PASS: TestMultiply/0*9=0 (0.00s)
	    --- PASS: TestMultiply/-7*6=-42 (0.00s)
	PASS
	ok      github.com/Vijayaraagavan/rolopos.git/menuconfig/buf    0.167s
	-------------------------------------------------------------------------
	

if any case fail, other cases will not be disturbed
	=== RUN   TestMultiply
	=== RUN   TestMultiply/2*3=6
	=== RUN   TestMultiply/10*5=50
	=== RUN   TestMultiply/-8*-3=24
	=== RUN   TestMultiply/0*9=1
	    Hello_test.go:26: got 0; want 1
	=== RUN   TestMultiply/-7*6=-42
	--- FAIL: TestMultiply (0.00s)
	    --- PASS: TestMultiply/2*3=6 (0.00s)
	    --- PASS: TestMultiply/10*5=50 (0.00s)
	    --- PASS: TestMultiply/-8*-3=24 (0.00s)
	    --- FAIL: TestMultiply/0*9=1 (0.00s)
	    --- PASS: TestMultiply/-7*6=-42 (0.00s)
	FAIL
	exit status 1
	FAIL    github.com/Vijayaraagavan/rolopos.git/menuconfig/buf    0.190s
	-------------------------------------------------------------------------
	

to run test or benchmark use -run or -bench flag
to run specific testcase provide regex expression 
	go test -run=TestMultiply/"^[0-5]\*[0-5]" -v
[0-5] means a single character can be between 0 and 5
^ to tell the match should start at beginning of word
\* to escape special characters. Because simpe '*' will mean [0-5] can have 0 or more times (1,11,243,152...)

	=== RUN   TestMultiply
	=== RUN   TestMultiply/2*3=6
	--- PASS: TestMultiply (0.00s)
	    --- PASS: TestMultiply/2*3=6 (0.00s)
	PASS
	ok      github.com/Vijayaraagavan/rolopos.git/menuconfig/buf    0.166s
	-------------------------------------------------------------------------

see fibonacci folder for benchmark
	


