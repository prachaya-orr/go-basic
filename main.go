package main

import "errors"

func genErr() error {
	return errors.New("some error")
}

func main() {
	// a := 2
	// if a != 2 {
	// 	fmt.Println("is not 2")
	// } else if a == 2 {
	// 	fmt.Println(a)
	// }

	//[80,100]->A
	//[70,80)->B
	//[60,70)->C
	//[50,60)->D
	//[0,50)->F

	// sc := 60
	// if sc >= 80 {
	// 	fmt.Println("A")
	// } else if sc >= 70 {
	// 	fmt.Println("B")
	// } else if sc >= 60 {
	// 	fmt.Println("C")
	// } else if sc >= 50 {
	// 	fmt.Println("D")
	// } else {
	// 	fmt.Println("F")
	// }

	if err := genErr(); err != nil {
		panic(err.Error())
	}

}
