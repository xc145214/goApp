package main
import "fmt"
//继承特性

type parent struct {
	val int
}

type child struct {
	parent
	num int
}

func main() {
	var c child

	c = child{parent{1}, 2}

	fmt.Println(c.val)
	fmt.Println(c.num)
}