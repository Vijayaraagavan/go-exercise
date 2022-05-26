package reg
import(
  "regexp"
  "fmt"
)
func main() {
  str := "seafood foo fool"
  re, _ := regexp.Compile(`foo.?`) // foo + any character + this character 1 times (0 or 1)
  match := re.FindAllString(str)
  fmt.Println(match)               // [food foo fool]
  /*
  when a '?' follows a repetition, it is greedy
  when a '??' follows a repetition, it is non-greedy
  so same str, if pattern is `foo.??` => o/p: [foo] (it prefers 0 number of times of .(any character) )
  
  if pattern is `fool?` => [foo foo fool] (it prefers more 'l' that is one times, if no 'l' its fine
  if pattern is `fool??` => [foo foo foo] (it prefers 0 'l')
  if str = "seafood fooll fool" and pattern =  `fool?` => [fool fool]
  */
  ------------------------
  
  str = "-a-abb-b"
  re, _ = regexp.Compile(`ab*`) // 0times b or many times b => a | ab| abbb....
  re.ReplaceAllString(str, "+")
  fmt.Println(str)
  -------------------------
  
  re, _ = regexp.Compile(`a`)
	fmt.Printf("%q\n", re.Split("banana", -1)) // ["b" "n" "n" ""]
	fmt.Printf("%q\n", re.Split("banana", 0))  // [] (nil slice)
	fmt.Printf("%q\n", re.Split("banana", 1))  // ["banana"]
	fmt.Printf("%q\n", re.Split("banana", 2))  // ["b" "nana"]
	
	re, _ = regexp.Compile(`z+`)
	fmt.Printf("%q\n", re.Split("pizza", -1)) // ["pi" "a"]
	fmt.Printf("%q\n", re.Split("pizza", 0))  // [] (nil slice)
	fmt.Printf("%q\n", re.Split("pizza", 1))  // ["pizza"]
	fmt.Printf("%q\n", re.Split("pizza", 2))  // ["pi" "a"]
  
  /*
  Use the Split method to slice a string into substrings separated by the regexp. 
  It returns a slice of the substrings between those expression matches.
  A return value of nil indicates no match.
  */
  ------------------------
  
  re := regexp.MustCompile(`foo(d?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`),-1))   //  [["food" "d"] ["foo" ""]]
  
  re := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`),-1))   // [["food" "d"] ["fool" "l"]]
  
  /*
  here (d?) is the subgroup. this func returns all matches and then it searches for any match of substring
  first it finds "food" => then it checks `(d?)` inside food and fool => prints d
  in second example the sub pattern is `(.?)` inside food and fool => prints d and l
}
