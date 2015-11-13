package main
import (
	"fmt"
	"time"
)

func fibonacci(c,quit chan int){
	x,y := 1,1
	for{
		select {
		case c <-x :
			x,y = y,x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func show2(c,quit chan int){
	for i:=0;i<10 ;i++  {
		fmt.Println(<- c)
	}
	quit <- 0
}

func main() {
	data := make(chan int)
	leave := make(chan int)

	go show2(data,leave)
	go fibonacci(data,leave)

	for  {
		time.Sleep(100)
	}
}