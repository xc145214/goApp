package main
import "fmt"

//类
type Integer int

//方法
func (a Integer)less(b Integer) bool {
	return a < b
}

func main() {

	fmt.Println("hello")
	var a,b Integer = 1,2
	fmt.Println(a,b)

	res := a.less(b)
	fmt.Println(res)

	//值传递
	var arrayA [3]int = [3]int{1,2,3}
	arrayB := arrayA
	arrayB[1]++
	fmt.Println(arrayA,arrayB)
}