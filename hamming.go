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
