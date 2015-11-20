package main
import "fmt"

//延迟方法 ：后进先出
func defun(){

	for i:=0;i<5;i++{
		defer fmt.Printf("%d ",i)
	}
}

//函数默认值
func f()(ret int)  {
	defer func() {
		ret ++
	}()

	return 0
}

//多参函数
func myfunc( arg ...int){
	fmt.Println(arg)
}



func main() {

	defun()


	i := f()
	fmt.Println(i)


	myfunc(1,2,3,4,4,5)


	//函数作为值
	a := func() { fmt.Println("hello world!") }

	a()
}
