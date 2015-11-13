package main

import "fmt"

//减法
func sub(a int,b int) int{
	return a-b;
}

//比较(if)
func compare(a int,b int) {
	if(a>b){
		fmt.Println(a,"is greater ")
	}else{
		fmt.Println(a,"is not greater")
	}
}

//switch
func chose(a int){
	switch (a) {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("error")
	}
}

//for
func showList(data int){
	var index int
	index =0
	for   {
		if(index >=data){
			break
		}
		fmt.Println(index)
		index ++
		continue
	}
}

func main() {
	fmt.Println(sub(2,3))

	compare(2,3)

	chose(3)

	showList(10)
}