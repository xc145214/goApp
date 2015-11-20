package main
import "fmt"


func main() {

	n := []int{5, -1, 0, 12, 3, 5}

	fmt.Printf("unsort %v\n", n)

	bubblesort(n)

	fmt.Printf("sort %v\n", n)
}


//冒泡排序
func bubblesort(n []int) {
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}

}




