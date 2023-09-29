package module

import "fmt"

func Sum(msg string, a, b int) (int, string) {
	fmt.Println(msg)
	return a + b, "\nHello: " + msg
}
