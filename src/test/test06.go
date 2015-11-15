package main
import (
	"fmt"
	"unicode/utf8"
)

//字符串统计
func getCode() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("String %s\n Length: %d,Runes: %d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
}

//字符串转换
func reverse() {
	s := "foobar"
	a := []rune(s)

	fmt.Printf("%s\n", string(a))

	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Printf("%s\n", string(a))
}


func main() {

	getCode()


	reverse()
}
