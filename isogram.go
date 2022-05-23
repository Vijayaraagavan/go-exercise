/*
An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

Examples of isograms:

lumberjacks
background
downstream
six-year-old
The word isograms, however, is not an isogram, because the s repeats.
*/

//implementing this problem using string could be tricky. Using slice or maps could be easy
package isogram
import (
    "strings"
    "unicode"
    )
func IsIsogram(word string) bool {
    var x string = ""
    for _, val := range word {
        if !unicode.IsLetter(val) {
            continue
        }
        if strings.Contains(strings.ToLower(x), strings.ToLower(string(val))) {
            return false
        } else {
         x += string(val)
        }
    }
	return true
	panic("Please implement the IsIsogram function")
}
// unicode is very useful in checking numbers, letters, special characters. It expects rune datatype and returns bool
// strings.Contains(s, substr string) bool { is also useful function for checking substr in s
--------------------------------------------------------------

// strings.ContainsRune() is useful when we give rune datatype as input instead string
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	for i, c := range s {
		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
			return false
		}
	}
	return true
}
