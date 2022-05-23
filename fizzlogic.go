/*
has 3 as a factor, add 'Pling' to the result.
has 5 as a factor, add 'Plang' to the result.
has 7 as a factor, add 'Plong' to the result.
does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
Fizz buzz logic
*/
package raindrops
import (
    "fmt"
    "strconv"
    )

func Convert(number int) string {
    hasThree := "Pling"
    hasFive := "Plang"
    hasSeven := "Plong"
    
    iif (number % 3 != 0) {
        hasThree = ""
    }
	  if( number % 5 != 0) {
    	hasFive = ""
    }
	  if (number % 7 != 0) {
    	hasSeven = ""
    } 
  res := strconv.FormatInt(int64(number), 10)     // or use strconv.Itoa()
	x := fmt.Sprintf("%s%s%s", hasThree, hasFive, hasSeven)
    if len(x) == 0 {
        return res
    }
    return x
	panic("Please implement the Convert function")
}
