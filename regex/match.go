package reg
import (
  "fmt"
  "regexp"
)

func main() {
  search := "All is fine"
  
  match1, err := regexp.MatchString("is", search)
  fmt.Println("Match found: ", match1, "| Pattern error: ", err)  // Match found: true | Pattern error: <nil>
  
  match2, err := regexp.MatchString("no", search)
  fmt.Println("Match found: ", match1, "Pattern error: ", err)  // Match found: false | Pattern error: <nil>
  
  match3, err := regexp.MatchString("is(s", search)
  fmt.Println("Match found: ", match1, "Pattern error: ", err)  // Match found: false | Pattern error: error parsing regexp: missing....
  
  ------------------------------------------------
  /*
  to store complicated regular expression for reuse later we use Compile() that parses regular expression and returns a 
  Regexp type
      func Compile(expr string) (*Regexp, error)
  Regexp is an struct and it has many receiving functions that we can use for regex
  */
  
  re, _ := regexp.Compile("i")
  search := "All is fine Allout attack"
  match := re.FindStringIndex(search)
  fmt.Println(match)  //[5, 9] it returns slice of int telling at index it starts and at which 
                      // index it ends in which the match is found
  
  ------------------------------------------------------
  match := re.FindAllStringSubmatchIndex(seach, -1) //[[0 3] [17 20]] returns slice of slice of int, all starting and ending
  fmt.Println(match)                                // index of all matches. -1 means unlimited. 2 means first 2 match
  
  ------------------------------------------------------
  re, _ := regexp.Compile(" ")
  match := re.ReplaceAllString(str, "+")
  fmt.Println(match)      //All+is+fine+Allout+attack
  
  ------------------------------------------------------
  search := "All I do"+ " is code everytime."
  re4, _ := regexp.Compile("[aeiou]+")
  match6 := re4.ReplaceAllStringFunc(search, strings.ToUpper) //each found match is passed as argument to given func 
  fmt.Println(match6)                                         //and return value is replaced
  
  -----------------------------------------------------
  
}
