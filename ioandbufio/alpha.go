package main

import(
  "fmt" 
  "strings"
  "io"

)
type alphaReader struct {
  reader io.Reader
}

func createAlpha (reader io.Reader) *alphaReader {
  return &alphaReader{reader: reader}
}

func alpha (r byte) byte {			//filters only small and capital letters
  if ( r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z' ) {
    return r
  }
  return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
  n, err := a.reader.Read(p)		  //Read() reads data until it reaches EOF error
  if err!=nil{						        //by default, Read() separates data by "space"
    return n, err					        //so in first iteration i.e first Read() call, it reads "hello"
  }									              // in second it reads "is"
  buf := make([]byte, n)
  for i:=0; i<n; i++ {
    if char := alpha(p[i]); char !=0 {	//loop runs 5 times ['h','e','l','l','o']
	  buf[i] = char
	  fmt.Println(string(char))			//prints h e l l o (each in new line)
    }
  }
 copy(p, buf)
 fmt.Println(p)
 return n, nil
}

func main(){
	reader := createAlpha(strings.NewReader("hello is  9 ma king"))
	p := make([]byte, 5)
	res := ""
	for {
		fmt.Println("func calling")
		n, err := reader.Read(p)		// in first Read() call it gets "hello"
		if err!=nil{
			if err ==io.EOF {
				break
			}
			fmt.Print(string(p[:n]))
			panic(err)
		}
		res += string(p)	// hello + is + ma + king
	}

	fmt.Println(res)	//helloismaking
}
                     
/*
 CHAINING READERS:
  It is common to use a reader as the source of another reader. This chaining of readers allows one reader to reuse logic
  from another reader which is declared as its field inside its type. Like strings.NewReader() returns an io.Reader object
  it is placed as an field in alpha struct type. 
  Then we implement Read() func to implement reader interface
  Inside Read() method of alpha, we use Read() of strings type. Since we are using io.Reader in func parameter while creating 
  alpha struct, we can accept any io.Reader type and apply same method
*/
