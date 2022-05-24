package fibrec
import (
  "fmt"
  "time"
)
func Fibrecursive(n int) int {
  if n < 2 {
    return n
  }
  left_node := Fibrecursive(n-1)
  right_node := Fibrecursive(n-2)
  
  result := left_node + right_node
  
  return result
}

func main() {
  var in = []int{10,20,30}
  var result []int
  var timer []time.Duration
  var sum_milli int64     // time.Duration is of type int64, it has Seconds(), Milliseconds, Minutes, and so on.... to return values in either float64/int64
  var sum_min float64     // we will use it to add these values to get total time taken
  for i:=0; i<len(in); i++ {
    start := time.Now()
    data := in[i]
    res := Fibrecursive(data)
    result = append(result, res)
    timer = append(timer, time.Since(start))
  }
  for _, i := range timer {
	  sum_milli += i.Milliseconds()
	  sum_min += i.Minutes()
  }
  for k, i := range result {
    fmt.Println( in[k], i, timer[k]) 
  }
  fmt.Println("It totally took :", sum_milli, "ms", sum_min, "minutes")

}

/*
10 55 1.442µs
20 6765 47.577µs
30 832040 5.786663ms
It will take approx 1m28s to calculate for input 50

this recursion is inefficient. Most of the time we should use dynamic programming to reduce time
*/
