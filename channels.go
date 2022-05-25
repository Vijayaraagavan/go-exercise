package chann
import (
  "fmt"
)

func main(){
  datapipe := make(chan int)      //channel can be any datatype, but it is homogeneous
  fmt.Println(<- datapipe)        // asking data from channel, if it is empty then it instructs main thread/program to stop further processing and wait
                                  // until data is poured into channel (deadLock)
                                //since the main process is blocked, the only way to fill data into channels is another go routine
}
-------------------------------
func main(){
  datapipe := make(chan int) 
  datapipe <- 5               //will this work?..No! As soon as one go routine fills data into the channel that goroutine is blocked until some other goroutine receives it
  fmt.Println( <-datapipe)    //again deadLock
}
-----------------------------

/*
This is because by default the channel has no storage capacity. That means when a message is sent it should be received immediately orelse it is stuck
with the sender.
When we scale our application there can be thousands of threads(sub process or branch of main thread(main program)) or go routines . In that case we can't 
block our thread for so long for each data transmission.
So we give channels a little storage to store messages (cache or buffer) and it is called buffered channels. Here the thread doesn't need to wait until
other thread receives or vice versa. Because the data is stored in buffer, the thread can continue its next work
If the buffer or storage is full, then the thread will be blocked
*/
func main() {
  datapipe := make(chan int, 3)
  datapipe <- 5
  datapipe <- 8
  datapipe <- 15
  fmt.Println( <-datapipe, <- datapipe, <- datapipe)
}
// if we add another data (exceeding channel length), this thread (main) is blocked and waits for other thread to receive it, so deadLock
// if you receive one data before adding 4th data, then still buffer size remains 3, no problem. So the point is at any time the channel buffer should 
// not store more than its size or the sender should not send more than its size before current data is extracted
