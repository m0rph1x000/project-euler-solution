package main

import "fmt"

var a, b int64 = 1, 1

func NextFibonacci() int64 {
	a, b = b, a+b
	return a
}

func main() {
	var total int64 = 0
	for i := 0; ; i++ {
		fib := NextFibonacci()
		if fib%2 == 0 {
			total += fib
		}
		if fib > 4_000_000 {
			break
		}
	}
	fmt.Println(total)
}
