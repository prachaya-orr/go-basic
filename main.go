package main

import "fmt"

// type Number interface {
// 	int | float64
// }

// func main() {
// 	numsInt := []int{1, 2, 3, 45}
// 	numsFloat64 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

// 	// fmt.Println(sumInt(numsInt))
// 	// fmt.Println(sumFloat64(numsFloat64))
// 	fmt.Println(sum(numsInt))
// 	fmt.Println(sum(numsFloat64))
// }

// func sum[V Number](nums []V) V {
// 	var sum V
// 	for _, num := range nums {
// 		sum += num
// 	}

// 	return sum
// }

// func sumInt(nums []int) int {
// 	// sum := 0
// 	var sum int
// 	for _, num := range nums {
// 		sum += num
// 	}

// 	return sum
// }

// func sumFloat64(nums []float64) float64 {
// 	// sum := 0.0
// 	var sum float64
// 	for _, num := range nums {
// 		sum += num
// 	}

// 	return sum
// }

type GameOrMovie interface {
	getPrice() int
}
type Game struct {
	Title    string
	Platform string
	Price    int
}

type Movie struct {
	Title string
	Price int
}

func (g Game) getPrice() int {
	return g.Price
}
func (g Movie) getPrice() int {
	return g.Price
}

func sum[V GameOrMovie](objs []V) int {
	var result int
	for _, obj := range objs {
		result += obj.getPrice()
	}
	return result
}

func main() {
	games := []Game{
		{
			Title:    "Minecraft",
			Platform: "PC",
			Price:    30,
		},
		{
			Title:    "The Sims 4",
			Platform: "PC",
			Price:    50,
		},
	}
	movies := []Movie{
		{
			Title: "Star war",
			Price: 10,
		},
		{
			Title: "Avenger",
			Price: 20,
		},
	}

	fmt.Println(sum(games))
	fmt.Println(sum(movies))
}
