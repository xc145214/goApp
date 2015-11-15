package main
import "fmt"

func printA() {
	for i := 0; i < 100; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
}

func stdPrintA(){
	str := "A"
	for i:=0;i<100 ;i++  {
		fmt.Printf("%s\n",str)
		str += "A"
	}
}

func main() {

	//printA()

	stdPrintA()


}