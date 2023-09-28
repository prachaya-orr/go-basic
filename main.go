package main

import "fmt"

func main() {
	// var name type = value
	// var a int = 2
	// var msg string = "Rock"

	var a any = "test"

	b, ok := a.(string)
	if !ok {
		panic("err")
	}
	fmt.Println(b)
}
