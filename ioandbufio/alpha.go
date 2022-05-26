package buf

import(
  "fmt" 
)
type alphaReader struct {
  reader io.Reader
}

func createAlpha (reader io.Reader) *alphaReader {
  return &alphaReader{reader: reader}
}

func alpha (r byte) byte {
  if ( r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' ) {
    return r
  }
  return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
  n, err := a.reader.Read(p)
  if err!=nil{
    return n, err
  }
  buf := make([]byte, n)
  for i:=0; i<n; i++ {
    if char := alpha(p[i]; char !=0 {
      buf[i] = char
    }
  }
 copy(p, buf)
 return n, nil
}
                     
                     
                     
