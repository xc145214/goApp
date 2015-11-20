package main
import "fmt"

//求平均值
func arvange(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0.0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return
}

func main() {

	//定义数组
	var xs = []float64{1.23,3.21,4.23,5.23,5.32}

	avg := arvange(xs)

	fmt.Printf("avg: %f",avg)
	
}
