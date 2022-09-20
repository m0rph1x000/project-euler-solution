package main

import (
	"fmt"
	"math"
)

func Eratosthenes(n int64) []int64 {
	res := make([]int64, 0, n/int64(math.Sqrt(float64(n))))
	sieve := make([]bool, n+1)
	z := int64(math.Sqrt(float64(n)))
	var i, j int64
	for i = 2; i <= z; i++ {
		if !sieve[i] {
			for j = i * i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}

	for i = 2; i <= n; i++ {
		if !sieve[i] {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	primes := Eratosthenes(2_000_000)
	var sum int64 = 0
	for _, v := range primes {
		sum += v
	}
	fmt.Println(sum)
}
