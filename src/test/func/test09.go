package main

var a = 6
func p() {
	println(a)
}

func q() {
	a = 5
	println(a)
}

func main() {

	p()
	q()
	p()
}