package main
import "fmt"



func main() {
	var array [100]int
	slice := array[0:99]

	slice[98] = 'a'
	fmt.Println(slice)

	s0 := []int{0, 0}
	s1 := append(s0, 2)
	s2 := append(s1, 3, 5, 7)
	s3 := append(s2, s0...)
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)



	//map
	monthdays := map[string]int{
		"Jan":31, "Feb":28, "Mar":31,
		"Apr":30, "May":31, "Jun":30,
		"Jul":31, "Aug":31, "Sep":30,
		"Oct":31, "Nov":30, "Dec":31,
	}
	fmt.Println(monthdays)

	year := 0
	for _, days := range monthdays {
		year += days
	}

	fmt.Println("Numbers of days in a year:%d\n",year)

}
