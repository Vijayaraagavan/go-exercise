package fibrec
import (
  "fmt"
  "time"
  "sync"
"runtime"
)
var result = make([]int,0)	// declare result and timer as global so that they can be easily accessed in functions
var timer = make([]time.Duration, 0) //size is 0, because we are going to append dynamically

func Fibrecursive(n int, check int, wg *sync.WaitGroup, in []int, index int, start time.Time) int {
	/*
	n - current input for fibonacci ( it can change dynamically, like 10, 10-1=9, 10-2=8,...)
	check - const data of input ( this is the first input before recursion starts, 10 which is constant)
	wg - to control go routines
	in - input slice (we can get only one input or multiple inputs)
	index - current index to find the input in input slice
	start - current time during execution of each go routines
  so why do we need this much parameters. because we are using go routines it is necessary to watch over them using either channels or waitgroup
  orelse main program will finish its execution and terminates. channels are useful when we need to send info among go routines. Here waitgroup is enough just
  to stall the main program
  	it is just like counter. Add(1) adds 1 count to its counter. to decrement it use Done(). wait() stalls the programs until its counter reaches zero
	In normal function we can use "defer wg.Done()". but here if you use this then it will be called recursively. 
so we need to execute this statement only during root func call. suppose input is 50, it is split into 49 and 48 for recursive call. In those sub function
calls n will not be 50, so if condition will only execute when all recursive calls finished and returns its output to first_node and right_node, then next line executes
	This is ok for single input, where we can directly give 50 as check
	but when we give multiple input through slice, we need to give corect input for if condition (i.e check or in[index])
	Now there is a problem
	if you just pass iteration index i to go function, you will notice that i is incremented from 0 to 1. if slice length itself is 1, then accessing 
slice will raise "range out of index" error 
	This is because when a function is ran as go routine it is detached from main programs flow. It runs independently. so right after you assign i 
value to index it will start next loop, where it will increment i value then checks for loop condition. Only at that time, go routine starts executing. so 
index value is now 1 and not 0
	so i have used -1 as initial value. I am not sure whether it is correct approach
	Orelse we can use check variable which has current fib input

	*/
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
	runtime.GOMAXPROCS(4)	// it is not recommended to use parallel execution if there is not so many cores in CPU
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
  var sum_milli int64	//check time package. each function returns either float64 or int64 when time converted to numbers
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
