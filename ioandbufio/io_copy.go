package buf

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  proverbs := new(bytes.Buffer)               // source: bytes.Buffer (both io.Writer & io.Reader type)
  proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")
  
  file, err := os.Create("./infotech.txt")    //destination: os.File (both io.Writer & io.Reader type)
  if err !=nil {
    panic(err)
  }
  defer file.Close()
  
  if _, err := io.Copy(file, proverbs); err != nil {  // copies proverbs to file
    panic(err)
  }
  fmt.Println("file created")
}

/*
without copy() how we can do this?
First create a io.Reader type eg: strings.NewReader or bytes.Buffer like any type
Second store data in it
use Read() func to convert data source => []byte
create a file which returns os.File type
use Write() func to store / write data to it from []byte => its data stream
------------------
but with io.Copy() this process is simplified
Now we can use any io.Reader in place of source, and any io.Writer type in place of dest
data source => Read() => []byte => Write() => data stream
*/
