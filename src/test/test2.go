package main
import "fmt"

//定义常量a,b为枚举
const (
	a = iota
	b
	c = 0
	d string = "0"
)

func main() {

	//字符串赋值
	s := "hello wrold!"
	fmt.Print(s)

	var s1 string = "hello"
	//s1[0]="c"		字符串无法修改
	fmt.Print(s1)

	sc := []rune(s1)
	fmt.Print(sc)
	sc[0] = 'c'
	s2 := string(sc)

	fmt.Print("%s\n", s2)

	//多行字符串

	ss := "starting part" +
	"ending part"

	fmt.Println(ss)

}
