package main

import "fmt"

func main() {

	// Basic -> sync.Waitgroup
	// a := []int{1, 2, 3, 4, 5}

	// var wg sync.WaitGroup
	// for i := range a {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Printf("%v ", a[i])
	// 	}(i)
	// }
	// wg.Wait()

	// var wg sync.WaitGroup
	// wg.Add(len(a))
	// go func() {
	// 	for i := range a {
	// 		defer wg.Done()
	// 		fmt.Printf("%v ", a[i])
	// 	}
	// }()
	// wg.Wait()

	// Chanel
	numberCh := make(chan int)
	msgCh := make(chan string)

	// number
	// chan<- int รับค่าได้อย่างเดียว ไม่ assign
	go func(numberCh chan<- int) {
		numberCh <- 10
	}(numberCh)

	// msg
	go func(msgCh chan<- string) {
		msgCh <- "hello world"
	}(msgCh)

	number := <-numberCh
	msg := <-msgCh

	fmt.Println(number)
	fmt.Println(msg)

	// Worker Pool

}
