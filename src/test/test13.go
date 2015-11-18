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




func main() {

	defun()


	i := f()
	fmt.Println(i)

}
