package fibrec
import (
 "fmt" 
)
func Fibdynamic(n int, cache map[int]int) int {
  if n < 2 {
    cache[n] = n
    return n
  }
  //first check if we already have stored value for input orelse compute and store it 
  if _, ok := cache[n-1]; !ok {
    cache[n-1] = Fibdynamic(n-1, cache)
  }
  
  if _, ok := cache[n-2]; !ok {
    cache[n-2] = Fibdynamic(n-2, cache) 
  }
  // each time we retrieve output from cache
  result = cache[n-1] + cache[n-2]
  return result
}

func main() {
  var in = []int{10,20,30}
  var result []int
  var timer []time.Duration   // time.Now() returns current time of type time.Time, whereas time.Since() returns total time from given start time of type 
                               // time.Duration (type is mostly struct with its own field and receiving function)
  
  var cache = make(map[int]int)   // always use make() func, and not map literal(var cache map[int]int), bcoz make returns dynamic array with default values, 
  for i:=0; i<len(in); i++ {      // so there will be no error when we store value in any index even if map is empty
    start := time.Now()
    data := in[i]
    res := Fibdynamic(data, cache)
    result = append(result, res)
    timer = append(timer, time.Since(start))
  }
  
  for k, i := range result {
    fmt.Println( in[k], i, timer[k]) 
  } 
}


/*
The output is 
10 55 9.231µs
20 6765 10.787µs
30 832040 11.707µs

Memoization: Not to be confused with Memorization, Memoization is a technique for improving the performance of a recursive function/algorithm. In other words, it offers an optimization to speed up programs by storing the solutions of expensive function calls and returning the cached solution when the same inputs are fed to the program again.
Tabulation: Also known as the Bottom-Up technique, this is another optimization technique used in dynamic [rogramming. As you might guess, this technique takes a tabular approach by first filling up a table, then solves the original problem based on the solutions filled in the table.

*/
