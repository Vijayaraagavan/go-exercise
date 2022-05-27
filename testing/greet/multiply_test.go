// table driven test
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
  /*
   "go test"     =>  will run all tests
   "go test -v"  => -v flag will print out the names of all the executed test functions and time spent for their execution
  
  t.Errorf() == first invoke t.Logf(), which logs text to the console either on test failures or when the -v flag is used
                Then it invokes t.Fail(), which marks the current function as failed without halting its execution
  t.Fatal() or t.Fatalf() => used to mark the current function as failed, stopping its execution immediatiely
            == t.Log() or t.Logf() followed by t.FailNow()
  Disadvantage : it cannot run specific test or testcase
                 t.Fatal() causes a test case to stop execution, then it does not run other testcases

*/
