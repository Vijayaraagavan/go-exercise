package buf
import (
  "fmt"
  "bufio"
)
type Writer int 
func (*Writer) Write(p []byte) ( n int, err error) {
  fmt.Println(len(p))
  return len(p), nil
}

func main() {
  fmt.Println("unbuffered I/O")
  w := new(Writer)
  w.Write([]byte{'a'})
  w.Write([]byte{'b'})
  w.Write([]byte{'c'})
  w.Write([]byte{'d'})      //output is 1 1 1 1 
  
  fmt.Println("buffered I/O")
  bw := bufio.NewWriterSize(w, 3)   // buffer of size 3 (default size is 4096 bytes)
  bw.Write([]byte{'a'})
  bw.Write([]byte{'b'})
  bw.Write([]byte{'c'})
  bw.Write([]byte{'d'})   //output is 3 1
  err := bw.Flush()       //we can flush output to dest using Flush() even if buffer is not full
  if err!=nil {
    panic(err)
  }
}

/*
package bufio helps with buffered I/O. So its obvious in its naming itself. buffer indicates it has some sort of storage
for input/output data. 
Doing many small writes can hurt performance. Each write is ultimatelly a syscall to operating system and if do it frequently
it can put a burden on cpu. 
Devices like disks work better with block-aligned data rather than a single data each time.
To avoid this problem go is shipped with bufio.Writer (implements io.Writer interface) as a replacement for io.Write 
Data instead going straight to destination it is stored in a buffer and when the buffer is full they are automatically sent
to the destination. Or we have func to flush data when we needed
          poducer => buffer => io.Writer
          
   producer         buffer           destination (io.Writer)
 
   a    ----->   a
   b    ----->   ab
   c    ----->   abc
   d    ----->   abcd
   e    ----->   e      ------>   abcd
   f    ----->   ef               abcd
   g    ----->   efg              abcd
   h    ----->   efgh             abcd
   i    ----->   i      ------>   abcdefgh
   
   bufio.Writer uses []byte as buffer storage
*/
