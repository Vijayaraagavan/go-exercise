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
func Source(filename string, out chan int, wg *sync.WaitGroup){
	f, err := os.OpenFile(filename, os.O_RDWR| os.O_CREATE | os.O_APPEND, 0777)
	if err!=nil {
		panic(err)
	}
	fmt.Println("opened success")
	// file, err := ioutil.ReadAll(f)
	// if err!=nil {
	// 	panic(err)
	// }
	// fmt.Println(string(file))
	b := bufio.NewReader(f)
	for {
		rd, err := b.ReadString('\n')
		if err!=nil {
			if err.Error() == "EOF" {
				wg.Done()
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
	defer wg.Done()
}

func Splitter(in, odd, even chan int, wg *sync.WaitGroup) {
	for val := range in {
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

func Sum(in, out chan int, wg *sync.WaitGroup){
	sum :=0
	for i := range in {
		sum+=i
	}
	out <- sum
	wg.Done()
}

func Merger(even, odd chan int, wg *sync.WaitGroup, resultFile string){
	f, err := os.Create(resultFile)
	if err!=nil {
		panic(err)
	}
	for i:=0; i<2; i++ {
		select {
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
	wg.Add(2)
	var wg2 sync.WaitGroup
	wg2.Add(4)

	var odd = make(chan int)
	var even = make(chan int)
	var out = make(chan int)
	var sumeven = make(chan int)
	var sumodd = make(chan int)

	go Source("input1.dat", out, &wg)
	go Source("input2.dat", out, &wg)

	go Splitter(out, odd, even, &wg2)

	go Sum(even, sumeven, &wg2)
	go Sum(odd, sumodd, &wg2)

	go Merger(sumeven, sumodd, &wg2, "res.txt")
	wg.Wait()
	close(out)
	wg2.Wait()
}
