package buf
import "testing"

type testcase struct {
  arg1 int
  arg2 int
  want int
}
func TestMultiply(t *testing.T) {
  cases := []testcase{        //to test mulitple cases for a function, define a struct and store input and outputs
    {2, 3, 6},                //using slice of struct we range over them
    {10, 5, 50},
    {-8, -3, 24},
    {0, 9, 0},
    {-7, 6, -42},
  }
  
  for _, val := range cases {
    got := Multiply(val.arg1, val.arg2)
    if val.want != got {
      t.Errorf("Expected '%d', but got '%d'\n", val.want, got) 
  }
}
  // "go test"     =>  will run all tests
  // "go test -v"  => -v flag will print out the names of all the executed test functions and time spent for their execution
  
