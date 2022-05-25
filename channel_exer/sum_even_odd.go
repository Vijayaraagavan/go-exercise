package chann

import (
	"fmt"
	"os"
	// "io/ioutil"
	"bufio"
	"strconv"
	"sync"
	"strings"
)
//func to open a file and read values line by line then pour it in channel
func Source(filename string, out chan int, wg *sync.WaitGroup){
	f, err := os.OpenFile(filename, os.O_RDWR| os.O_CREATE | os.O_APPEND, 0777) //these three flags are very important
	if err!=nil {					//os.O_RDWR gives read/write access, if there is no file, os.O_CREATE creates new file
		panic(err)				// it uses 0777 code as file permission for newly created file
	}
	fmt.Println("opened success")
	// file, err := ioutil.ReadAll(f)	this is io package's func (see bufio go file for reader interface)
	// if err!=nil {
	// 	panic(err)
	// }
	// fmt.Println(string(file))
	b := bufio.NewReader(f)			//creates a var with type bufio.Reader
	for {
		rd, err := b.ReadString('\n')	//bufio.Reader's receiver function. reads a string until it sees delimiter
		if err!=nil {
			if err.Error() == "EOF" {	//after last input in input1.dat file make sure you press enter(newline-\n) 
				wg.Done()		//orelse last input will be ignored
				return
			} else {
				panic(err)
			}
		}
		str := strings.ReplaceAll(rd, "\n", "")
		i, err := strconv.Atoi(str)
		if err!=nil {
			panic(err)
		}
		out <- i
		fmt.Println("done")
	}
	fmt.Println("exiting")
}

//After getting all inputs from different files, split values into even and odd, then send it to 
//corresponding channel(even and odd)
func Splitter(in, odd, even chan int, wg *sync.WaitGroup) {
	for val := range in {		//range can be used to iterate channels
		fmt.Println("enter")
		switch val%2 {
		case 0: 
			even <- val
		case 1:
			odd <- val
		}
	}
	fmt.Println("exit split")
	close(even)
	close(odd)
	wg.Done()
}

//func to calculate sum (same for odd/even) and send result to given out channel
func Sum(in, out chan int, wg *sync.WaitGroup){
	sum :=0
	for i := range in {
		sum+=i
	}
	out <- sum
	wg.Done()
}

//func to store result from even and odd channel into new text file
func Merger(even, odd chan int, wg *sync.WaitGroup, resultFile string){
	f, err := os.Create(resultFile)
	if err!=nil {
		panic(err)
	}
	for i:=0; i<2; i++ {		//we don't know which channel first gives input, it can be odd first or even first
		select {		//for either case prepare case statement
		case i:= <-even:
			f.Write([]byte(fmt.Sprintf("Even: %d\n", i)))
		case i:= <-odd:
			f.Write([]byte(fmt.Sprintf("Odd: %d\n", i)))
		}
		}
		wg.Done()
	}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)		// for main process
	var wg2 sync.WaitGroup
	wg2.Add(4)		// for sub process

	var odd = make(chan int)
	var even = make(chan int)
	var out = make(chan int)
	var sumeven = make(chan int)
	var sumodd = make(chan int)

	go Source("input1.dat", out, &wg)	//note that we are using go routine
	go Source("input2.dat", out, &wg)	//so opening and sending data in two files is done concurrently

	go Splitter(out, odd, even, &wg2)	//since all func are called as go routine, every func is blocked
						//and unblocked periodically depending on channels						
	go Sum(even, sumeven, &wg2)
	go Sum(odd, sumodd, &wg2)

	go Merger(sumeven, sumodd, &wg2, "res.txt")
	wg.Wait()
	close(out)
	wg2.Wait()
}
