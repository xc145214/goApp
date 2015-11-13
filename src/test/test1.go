package main

/**
 * 定义多个变量
 */
var x, y, z int
var s, n = "abc", 123

var (
	a int
	b float32
)



func main() {
	n, s := 0x1234, "hello world"
	println(x, s, n)

	//多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100
	println(i, data[i], data[0])

	//特殊只写变量 "_"，用于忽略值占位
	_,s = test()
	println(s)

	s1 := "abc"
	println(&s1)

	s1, y := "hello", 20 // 重新赋值: 与前 s 在同一层次的代码块中，且有新的变量被定义。
	println(&s1, y) // 通常函数多返回值 err 会被重复使用。

	{
		s2, z := 1000, 30 // 定义新同名变量: 不在同一层次代码块。
		println(&s2, z)
	}
}


func test()(int,string){
	return 1,"abc"
}
