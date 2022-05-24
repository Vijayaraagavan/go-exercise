package fibrec
import (
  "fmt"
  "time"
  "sync"
"runtime"
)
var result = make([]int,0)
var timer = make([]time.Duration, 0)

func Fibrecursive(n int, check int, wg *sync.WaitGroup, in []int, index int, start time.Time) int {
	// fmt.Println(" here i is ", index)
  if n < 2 {
    return n
  }
  left_node := Fibrecursive(n-1, check, wg, in, index, start)
  right_node := Fibrecursive(n-2, check, wg, in, index, start)
  
  res := left_node + right_node
  if n == in[index] {
	result = append(result, res)
    timer = append(timer, time.Since(start))
	wg.Done()
  }
  
  return res
}

func main() {
	runtime.GOMAXPROCS(4)
  var in = []int{10,20, 30, 40,
	10,20, 30, 40,
	10,20, 30, 40,
	10,20, 30, 40,
	10,20, 30, 40,
	10,20, 30, 40,
	10,20, 30, 40,
	10,20, 30, 40,
	10,20, 30, 40,
	10,20, 30, 40,
	10,20, 30, 40,
}
  var wg sync.WaitGroup
  var sum_milli int64
  var sum_min float64
  total_time := time.Now()

  for i:=0; i<len(in); i++ {
	index := -1
	wg.Add(1)
	data := in[i]
	index=i
	check := data
    start := time.Now()
	go func ()  {
		fmt.Println("go routine index ", index)
		Fibrecursive(data, check, &wg, in, index, start)
		fmt.Println("go routine over ", index)
	}()
}
wg.Wait()

total_time_end := time.Since(total_time)
  for _, i := range timer {
	  sum_milli += i.Milliseconds()
	  sum_min += i.Minutes()
  }
  fmt.Println(result)
  for k, i := range result {
	fmt.Println( in[k], i, timer[k])
}
fmt.Println("sum of all time taken by individuals :", sum_milli, "in ms", sum_min, "in minutes")
fmt.Println("total time: ", total_time_end)
}
