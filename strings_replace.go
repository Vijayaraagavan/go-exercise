/*

Two-fer or 2-fer is short for two for one. One for you and one for me.

Given a name, return a string with the message:

One for name, one for me.
Where "name" is the given name.

However, if the name is missing, return the string:

One for you, one for me.
Here are some examples:

Name	String to return
Alice	One for Alice, one for me.
Bob	One for Bob, one for me.
One for you, one for me.
Zaphod	One for Zaphod, one for me.

*/

package twofer
import "strings"
const pattern = "One for you, one for me."
// ShareWith tells you how to share with the given name
func ShareWith(name string) string {
	result := pattern
	if name != "" {
		result = strings.Replace(result, "you", name, -1)
	}
	return result
}

or

func ShareWith(name string) string {
    if name=="" {
        name = "you"
    }
	return fmt.Sprintf("One for %s, one for me.", name)
}

First solution is best to use. say we have n number of lines of paragraph, it is easy to replace all matching words using Replace() function. 
Also we can control the number the times the substitution can occur. (-1) means unlimited subsitution.

Second solution is suitable for few lines. As we know Printf(), Println() writes result to standard output(console), whereas Sprintf() returns output to function
FPrintf() mostly used in response.

