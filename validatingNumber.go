package luhn
import (
    "strconv"
    "fmt"
    "unicode"
)
func Valid(id string) bool {
    temp :=""

    // removing spaces 
	for _,k:= range id {
        if !unicode.IsDigit(k) && !unicode.IsSpace(k) {
			return false
		}
		if(unicode.IsSpace(k)){
			continue
		}
		temp += string(k)
	}
	// --------------
    //a string can't be less than 2 in length. Because if it is single digit, it can never be divisible by 10
    //even if you square it
	l := len(temp)
	fmt.Println(l)
    if l<2{
        return false
    }

    x := 0
    res := ""	
    // We are looping the string in reverse order with -2 as increment
    for i:=l-2; i>=0; i-=2 {
        x,_ = strconv.Atoi(string(temp[i]))
        x *= 2
        if x > 9 {
            x -=9
        }

    res += strconv.Itoa(x)			// Itoa() converts integer to string
    res += string(temp[i+1])		//since we are looping at 2 increment, we need to manually add in between values
	if (l%2)!=0 {					//if string has odd length("059"), only 9 is manually added and 0 is ignored, bcoz 
        res += string(temp[0])		// next loop won't run (condition i>=0)
	}
    }
        sum := 0
	for _, i := range res {
   	 val,_ := strconv.Atoi(string(i))
        sum += val
    }
	if sum % 10!=0 {
        return false
    } else {
    return true
    }
	panic("Please implement the Valid function")
}
//--------------------------------
/*
if we need to find second digit in every word separated by space, then use strings.Split(s, sep string) []string { to split them, and then range over them
temp := strings.Split(id, " ")
*/
//-----------------------------------

/* PROBLEM
Strings of length 1 or less are not valid. Spaces are allowed in the input, but they should be stripped before checking. 
All other non-digit characters are disallowed.

Example 1: valid credit card number
4539 3195 0343 6467
The first step of the Luhn algorithm is to double every second digit, starting from the right. We will be doubling

4_3_ 3_9_ 0_4_ 6_6_
If doubling the number results in a number greater than 9 then subtract 9 from the product. The results of our doubling:

8569 6195 0383 3437
Then sum all of the digits:

8+5+6+9+6+1+9+5+0+3+8+3+3+4+3+7 = 80
If the sum is evenly divisible by 10, then the number is valid. This number is valid!
*/
