package main
import "fmt"

//多态
type  act interface {
	write()
}

type xiaoming struct {


}

type xiaofang struct {


}

func (xm *xiaoming) write() {
	fmt.Println("xiaoming write")
}

func (xf *xiaofang)write() {
	fmt.Println("xiaofang write")
}

func main() {
	var w act

	xm := xiaoming{}
	xf := xiaofang{}

	w = &xm
	w.write()

	w = &xf
	w.write()
}