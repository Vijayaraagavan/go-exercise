package buf

type Reader interface {
  Read(p []byte) (n int, err error)
}
/*
This is an interface found in io package. To use receiver functions of type Reader interface "a datatype" must 
implement this interface. Implementing interface means just defining Read() in their type
eg: func (str string) Read(x []byte) (n int, err error) {
          /...do anything.../
     }
With this receiver function we have implemented Reader interface
Here we pass in a slice of bytes and ask Read() to fill it with its data - which it does until it runs out of data
It returns no of bytes read or error if something wrong. Additionally, if it has finished reading, it will return a 
special marker error called 'io.EOF' (End of File)
There are many types of Reader types available (types of reader types means many datatypes or types that implement io.Reader interface)
    1. Files
    2. string
    3. http.Request
    4. bytes.Buffer
*/
var r io.Reader       //if you use io.Reader interface as type, then it can accept any value that implements reader interface
var err error
r, err = os.Open("file.txt")  //before this, r is of type reader interface. Now it is os.File type which implements reader interface

---------------
var r io.Reader
r = strings.NewReader("read will return these bytes") // NewReader() takes a string and returns *io.Reader type that has slice of bytes of this string

-----------------
r = request.Body
----------------
var buf bytes.Buffer
r = &buf
----------------

//Using Readers
p := make([]bytes, 256)
n, err := r.Read(p)       // here we read from them directly (least used)
-------------------------
b, err := ioutil.ReadAll(r)   // reads everything from reader, and returns raw []byte data
-------------------------
n, err := io.Copy(w, r) // here w is an io.Writer type. Copy() reads all bytes from io.Reader and writes to an var of io.Writer type
--------------------------
err := json.NewDecoder(r).Decode(v)

----------------------------

/*
Mostly try to take an var of type io.Reader instead of string, file, http.request,... as argument in func
*/
func Reverse(s string) (sring, error)
 
//can be written as 
func Reverse(s io.Reader) io.Reader 

//now to use string as argument
r = Reverse(strings.NewReader("make this state"))
//or any reader type
f, err := os.Open("file.txt")
r = Reverse(f)

//Actually, this is an application of interface
            
