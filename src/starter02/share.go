package main
import (
	"fmt"
	"time"
)


func show() {
	for {
		fmt.Println("child")
		time.Sleep(1000)
	}
}

func main() {
	go show()

	for   {
		fmt.Println("parent")
		time.Sleep(1000)
	}
}