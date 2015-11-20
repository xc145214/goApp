package main

func order(a int, b int) (int, int) {

	if (a > b) {
		return b, a
	}
	return a, b
}

func main() {

	a, b := order(7, 2)
	print(a, b)

}
