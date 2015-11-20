package main
import "fmt"

// 斐波那契
func fibonacci(value int) [] int {
	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n - 1] + x[n - 2]
	}
	return x

}

func main() {

	for _, term := range fibonacci(10) {
		fmt.Printf("%v ", term)
	}

}
