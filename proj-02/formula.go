package main

import (
	"fmt"
	"math"
)

var p = math.Pow

func Fibonacci(n float64) float64 {
	param1 := (1 / math.Sqrt(5)) * p(((1+math.Sqrt(5))/2), n)
	param2 := (1 / math.Sqrt(5)) * p(((1-math.Sqrt(5))/2), n)
	formula := param1 - param2

	return formula
}

func main() {
	var total int64 = 0
	for i := float64(2); ; i++ {
		f := int64(math.Round(Fibonacci(i)))
		if f%2 == 0 {
			total += f
		}
		if f > 4_000_000 {
			break
		}
	}
	fmt.Println("Total: ", total)
}
