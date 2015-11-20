package main
import (
	"fmt"
	"strconv"
)
type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *stack)pop() int {
	s.i--
	return s.data[s.i]
}


func (s stack) string() string {
	var str string
	for i := 0; i <= s.i; i++ {
		str = str + "[" +
		strconv.Itoa(i) + "：" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main() {

	var s stack
	s.push(25)
	fmt.Printf("stack %v\n", s.string())
	s.push(15)
	fmt.Printf("stack %v\n", s.string())
	fmt.Printf("stack %v\n", s.pop())


}
