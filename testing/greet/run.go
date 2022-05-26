package main

import "fmt"

func main() {
  greet := Hello("vijay")   //no need any import statements, we can use exported func directly in another file within same package
	fmt.Println(greet) 
}
