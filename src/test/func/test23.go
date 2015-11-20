package main
import "fmt"


func main() {

	p2 := plusTwo()
	fmt.Printf("%v\n", p2(2))

	px := plusX(5)
	fmt.Printf("%v\n", px(6))
}

//返回函数
func plusTwo() func(i int) int {
	return func(x int) int {
		return x + 2
	}
}


//返回函数
func plusX(x int) func(int) int {
	return func(y int) int {return x + y}
}

