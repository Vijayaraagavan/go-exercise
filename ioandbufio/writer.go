package buf

import (
  "fmt"
  "os"
)

func main(){
  f, err := os.OpenFile("sample.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)  // returns os.File type 
  if err!=nil {                                 //os.File type already implements Reader interface (i.e Read())
    panic(err)
  }
  defer f.Close()
  
  n, err := f.Write([]byte("writing some info into a file"))  // os.File type also implements writer interface
  if err!=nil{                                                // so we can use Write() method
    panic(err)
  }
  fmt.Printf("written %d bytes", n)
}

----------------------------------------------------
type Writer interface {
  Write(p []byte) (n int, err error)
}

/*
Reader interface has Read() func, while Writer interface has Write() func

  in reader -> Data source => buffer ([]byte) => stores in given []byte variable ( file/string/request/ => bytes)
  in writer -> []byte data => data stream  (bytes => file/string/request)
  
  what read does is it take an data source and converts it to a raw slice of bytes of data and stores it in given []byte variable
  on the other hand, write takes slice of bytes and then it coverts it to some other data stream
  
  These are standard libraries (os.File, http.Request, json,..) they have already implemented some interface and all we ned
  to do is just utilise those methods/functions
  for example we don't need to define write() func for os.File types. we can directly call it. But when we create a struct or of some other
  data type as a type ("type gum struct"), we need to define Write() to implement writer interface and also has to decide
  what the Write() function should do. We need to follow only few rules like return type and values
*/

-----------------------------------------------------------
import (
  "fmt"
  "encoding/json"
  "bytes"
)
type user struct {
  Name string `json:"name"`
  Age int `json:"age"`
}

func main() {
  buf := new(bytes.Buffer)
  u := user{
    Name: "bob",
    Age: 20,
  }
  err := json.NewEncoder(buf).Encode(u)
  if err!=nil{
    panic(err)
  }
  fmt.Println(buf.String())
}

/*
  func NewEncoder(w io.Writer) *Encoder
  this func takes any type that implements io.Writer interface and returns address of Encoder type
  
  type Encoder struct {
  w io.Writer
  err error
  escapeHTML bool
  .
  .
  }
 
Now lets see what Encode() does
Encode() takes any data type as parameter (like interface[] accepts any type)
It has lot of functions inside. but in simple terms, it encodes the given type using marshal() func
Then is has to send the encoded value to some outuput. Here buf is the output. It can be a file or standard o/p (console), i.e io.writer object
Now Encode() utilitzes io.writer's write() func to send bytes of data to its respective data stream
for example, if buf is os.File type, then inside Encode definition it calls buf.Write( encoded bytes )
*/
