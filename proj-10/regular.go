package main

import "fmt"

func isPrime(p uint) bool {
	if p%2 == 0 {
		return false
	}
	for i := uint(3); i*i <= p; i += 2 {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var total uint = 0
	for i := uint(3); i < 2_000_000; i += 2 {
		if isPrime(i) {
			total += i
		}
	}
	fmt.Println(total + 2)
}
