package main
import "fmt"

//封装特性

type data struct {
	val int
}

func (p_data* data)set(num int) {
	p_data.val = num
}

func (p_data* data)show(){
	fmt.Println(p_data.val)
}

func main() {
	p_data := &data{4}
	p_data.set(5)
	p_data.show()
}