package chann

import (
  "fmt"
)
func show1(c chan int) {
	for i:=0; i<25; i++ {
		c <- i
	}
}
func show2(c chan int) {
	for i:=26; i<50; i++ {
		c <- i
	}
}
func show3(c chan int) {
	for i:=51; i<75; i++ {
		c <- i
	}
}
func main() {
  datapipe := make(chan int)

	go show1(datapipe)
	go show2(datapipe)
	go show3(datapipe)
	for i:=0; i< 50; i++ {
		fmt.Println(<- datapipe)
	}
}
/*
generating numbers from 0 to 75 is split by 3 go routines. Though this level of computation load will not make any changes. Think if the go routines 
actually handling large load
there is only one channel and it is sent as argument to all threads. When a goroutine / thread sends a number to  c channel, that thread is blocked until
main program or other go routine receives it. Other threads are also waiting to send data. Here we are not using any go routine to receive data, only send data.
Note that in main thread we are extracting only 50 data from channel. Does this means that the goroutine that sends data which main program doesn't receive
it? Yes!. so indicates show go routine is blocked indefinitely (deadLock). 
but it is no problem. because however main program/thread will end. if main thread is terminated, its sub thread will also be terminated
*/
