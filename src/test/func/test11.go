package main


func calc(a int, b int) (sum int,min int) {
	return a + b, a - b
}

func main() {
	a, b := calc(5, 7)

	print(a, b)
}
