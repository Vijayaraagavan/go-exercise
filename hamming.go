/*
When cells divide, their DNA replicates too. Sometimes during this process mistakes happen and single pieces of DNA get encoded with the incorrect information. If we compare two strands of DNA and count the differences between them we can see how many mistakes occurred. This is known as the "Hamming Distance".

We read DNA using the letters C,A,G and T. Two strands might look like this:

GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
^ ^ ^  ^ ^    ^^
They have 7 differences, and therefore the Hamming Distance is 7.
*/
package hamming
import "errors"
// Distance returns the hamming distance of two strings. The comparison is done
// at the character level, not the underlying bits.
func Distance(as, bs string) (int, error) {
	distance := 0
	if len(as) != len(bs) {
		return -1, errors.New("inputs must have the same length")
	}
	for i := range as {
		if bs[i] != as[i] {
			distance++
		}
	}
	return distance, nil
}

/*
use errors package for generating erros of type 'error'. If range has only one LHS variable it returns index. Orelse index and value ( key, value := range a { )
*/

func Distance(strand1 string, strand2 string) (distance int, e error) {
	if utf8.RuneCountInString(strand1) != utf8.RuneCountInString(strand2) {
		return 0, errors.New("strands must be of equal length. they're not. no soup for you")
	}
	// compare runes, not bytes
	strand1Runes := []rune(strand1)
	strand2Runes := []rune(strand2)
	// range iterates runes
	// but the index from range on a string increments to the position of the first byte of the next rune in the string,
	// not to the rune's position in a slice of runes, so we'd  need to manually increment the index ourselves.
	// iterating on strand1Runes := []rune(strand1) returns the slice index.
	for idx, rooney := range strand1Runes {
		if rooney != strand2Runes[idx] {
			distance++
		}
	}
	return distance, nil
}

/*
byte datatye represents only 127 characters (1byte). so if characters have utf8 encoding indexing will not be proper. Also len() function also returns wrong value.
Thats why we use rune datatype. it has length of 4bytes. for example chinese letters occupy 4bytes of data(thats what its unicode is). In that case len() will
return 4 for single character, but we expect 1. utf8.RuneCountInStrings() returns correct length
*/
