package learn

import "fmt"

type iTest interface {
	SayHello()
	Say(s string)
	Increment()
	GetValue() int
}

type iTestImpl struct {
	value int
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

	// var a iTest
	// a = &iTestImpl{}
	// a.SayHello()
	// a.Increment()
	// a.Increment()

	// fmt.Println(a.GetValue())

	a := 1

	fmt.Println(a)

}

// func add(a *int) {
// 	defer func() {
// 		fmt.Println("hi")
// 	}()
// 	*a++
// }

func (tst iTestImpl) SayHello() {
	fmt.Println("hello")
}

func (tst iTestImpl) Say(s string) {
	fmt.Println(s)
}

func (tst *iTestImpl) Increment() {
	tst.value++
}

func (tst *iTestImpl) GetValue() int {
	return tst.value
}
