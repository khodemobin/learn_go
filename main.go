package main

import "fmt"

type test struct {
	Name string
	Age  int
}

func main() {
	// var a [5]int = [5]int{}
	// a := []int{}
	// a := make([]int, 2)
	// a[0] = 1
	// a[1] = 2
	// a = append(a, 3, 4, 5, 6)

	// var a map[string]string = map[string]string{
	// 	"1": "1",
	// 	"2": "2",
	// }

	// a := make(map[string]string)
	// a["1"] = "1"
	// a["2"] = "2"

	// a := test{
	// 	Name: "mobin",
	// 	Age:  1,
	// }

	// a := 5
	// add(&a)

	a := 1
	fmt.Println(a)
}

// func add(a *int) {
// 	defer func() {
// 		fmt.Println("hi")
// 	}()
// 	*a++
// }
