package main
import "fmt"
type stack struct  {
	i int
	data [10]int
}

//错误的值传递
func (s stack) push1(k int){
	if s.i+1>9{
		return
	}
	s.data[s.i] = k
	s.i++
}



func main() {

	var s stack
	s.push1(25)
	fmt.Printf("stack %v\n",s)
	s.push1(15)
	fmt.Printf("stack %v\n",s)

}
