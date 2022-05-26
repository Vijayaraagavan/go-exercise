package main

import "fmt"

func Greet(user string) {     // naming must start with capital letter when you export a func or variable
  if (len(user) == 0) {
		return "hello dude"
	} else {
		return fmt.Sprintf("Hello %v!", user)
	} 
}
