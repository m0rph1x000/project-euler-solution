package main

import "fmt"

func main() {
	var a, b, c int

	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000-a; b++ {
			c = 1000 - a - b
			if (a*a)+(b*b) == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
