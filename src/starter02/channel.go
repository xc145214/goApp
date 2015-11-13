package main
import "fmt"

func show1(c chan int) {
	for {
		data := <-c
		if 1 == data {
			fmt.Println("receive!")
		}
	}
}

func main() {

	c := make(chan int)
	go show1(c)
	for {
		c <- 1
		fmt.Print("send ")
	}
}
