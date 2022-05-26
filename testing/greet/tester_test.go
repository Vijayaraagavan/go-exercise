package main

import "testing"

func TestWord (t *testing.T) {      //func name should start with "Test" and starting letter of test function
  expectation := "Hello vijay!"     // should be capital - "Word" so finally it should be TestWord
	actual := Hello("vijay")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestLength (t *testing.T) {
  expectation := 11
	actual := len(Hello("v"))
	if actual < expectation {
		t.Errorf("given length %v is less than expected %v", actual, expectation)
	}
}

/*
 run "go test" in cmd
 
output: 
    PASS
    ok      github.com/Vijayaraagavan/rolopos.git/menuconfig/buf    0.293s
    
if test cases fail
    --- FAIL: TestName (0.00s)
    Hello_test.go:9: Expected Hello vijay! but got Hello Vijay!
    --- FAIL: TestLenth (0.00s)
        Hello_test.go:17: given length 8 is less than expected 11
    FAIL
    exit status 1
    FAIL    github.com/Vijayaraagavan/rolopos.git/menuconfig/buf    0.366s
*/
