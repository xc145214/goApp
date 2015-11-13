package main

import "fmt"

func add(a int, b int) (c int) {
	c = a + b
	return c
}

//main方法
func main()  {
	c := add(1,2)
	fmt.Println(c)

	fmt.Println("hello world!")
}