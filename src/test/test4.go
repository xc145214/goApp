//fizzbuzz
package main
import (
	"fmt"
)

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if (i % 3 == 0 && i % 5 == 0) {
			fmt.Printf("FizzBuzz")
		}else if (i % 3 == 0) {
			fmt.Printf("Fizz")
		}else if (i % 5 == 0) {
			fmt.Printf("Buzz")
		}else {
			fmt.Printf("%v",i)
		}
		fmt.Println()
	}
}



func stdFizzBuzz() {
	const (
		FIZZ = 3
		BUZZ = 5
	)
	var p bool

	for i := 1; i < 100; i++ {
		p = false
		if i % FIZZ == 0 {
			fmt.Printf("Fizz")
			p = true
		}
		if i % BUZZ == 0 {
			fmt.Printf("Buzz")
			p = true
		}
		if !p {
			fmt.Printf("%v",i)
		}
		fmt.Println()
	}
}
func main() {

	fizzbuzz()

	stdFizzBuzz()
}