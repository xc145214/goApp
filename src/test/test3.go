package main
import "fmt"

//for循环
func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

//goto
func loop2() {
	i := 0
	Loop:
	fmt.Println(i)
	i++
	if (i < 10) {
		goto Loop
	}

}

//for循环
func loop3() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println("%v", arr)

	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("%v", a)
}




func main() {
	println("loop1")
	loop1()

	print("loop2")
	loop2()

	print("loop3")
	loop3()

}
