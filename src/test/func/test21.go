package main
import "fmt"

func max(l []int) (max int) {
	max = l[0]
	for _, v := range l {
		if (v > max) {
			max = v
		}
	}
	return
}

func min(l []int) (min int) {
	min = l[0]
	for _, v := range l {
		if (v < min) {
			min = v
		}
	}
	return
}


func maxAndMin(l []int) (max int,min int) {
	max = l[0]
	min = l[0]
	for _, v := range l {
		if (v > max) {
			max = v
		}
		if (v < min) {
			min = v
		}
	}
	return max, min
}

func main() {


	arr := []int{1, 4, 5, 3, 2, 7}

	print(max(arr))

	print(min(arr))


	fmt.Print(maxAndMin(arr))
}
