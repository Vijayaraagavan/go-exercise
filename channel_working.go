// DATA SENT/RECEIVED IS DONE ONE BY ONE IN UNBUFFERED CHANNEL
package chann

import(
  "fmt"
  "time"
)
func show(c chan int) {
  for i:=0; i<25; i++ {       //even though we use forloop, the fact is not all 25 inputs are given to channel
    c <- i                   // 0 inputted to c, now show() is blocked, in main() it is received, then show() resumes -> prints "data poured...."
    fmt.Println("data poured:", i)
  }
}
func main(){
  datapipe := make(chan int)
  for i:=0; i<2; i++ {      // here loop runs twice, so again it receives second data(1) from channel
    fmt.Println(<-datapipe) 
  }
  fmt.Println("main() ends")
}

/*
OUTPUT:
data poured:  0
0
1
---------
you will expect to print "data poured: 1" to print because after receiving second data, show() resumes
Thats right. but main() terminates before go routine resumes because loop ends and statement to execute further (except last println)
If you need to see the second print(data poured: 1), then add "time.Sleep(1 * time.Second)" to halt main() for 1 second. during that time go routine prints
it. but then during third data input show() is blocked indefinitely

OUTPUT:
data poured:  0
0
1
data poured:  1
---------------
*/
datapipe := make(chan int, 10)
/*
now that we added buffer of size 10, show() will not be blocked for 10 inputs. so show() sends 10 inputs continuously and prints "data pour.." independent
of whether other go routines receives it

OUTPUT:
data poured:  0
data poured:  1
data poured:  2
data poured:  3
data poured:  4
data poured:  5
data poured:  6
data poured:  7
data poured:  8
data poured:  9
data poured:  10
0
1
data poured:  11
*/
