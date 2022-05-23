/*
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
Examples
"cabbage" should be scored as worth 14 points:

3 points for C
1 point for A, twice
3 points for B, twice
2 points for G
1 point for E
*/
package scrabble
import (
	"fmt"
    "bytes"
    )
    
func Score(word string) int {
    words := bytes.ToLower([]byte(word))
    total := 0
    var store = map[byte]int{
        'a': 1,'e': 1,'i': 1,'o': 1,'u': 1,'l':1,'n': 1,'r': 1,'s': 1,'t': 1,
    	'd': 2,'g': 2,
    	'b': 3,'c': 3,'m': 3,'p': 3,
        'f': 4, 'h': 4, 'v': 4, 'w': 4,'y': 4,
        'k': 5,
        'j': 8, 'x': 8,
        'q': 10, 'z': 10,
    }
	for _, val := range words {
        total += store[val]
    }
	fmt.Println(total)
	return total
  panic("Please implement the Score function")
}

//switch is suitable when more variables point to the same value
func scoreOf(c rune) int {
    switch c {
    case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
        return 1
    case 'd', 'g':
        return 2
    case 'b', 'c', 'm', 'p':
        return 3
    case 'f', 'h', 'v', 'w', 'y':
        return 4
    case 'k':
        return 5
    case 'j', 'x':
        return 8
    case 'q', 'z':
        return 10
    default:
        return 0
    }
}
