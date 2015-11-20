package main
import "fmt"

func partthem(numbers ... int){
	for _,d := range numbers{
		fmt.Printf("%d\n",d)
	}
}
func main() {

	partthem(1,2,3,6,4,5)
	partthem(1,2,3,67,6,4,4,5)

}
