package main

import (
	"fmt"
	"math"
)

func main() {
	num := 600851475143
	upperLimit := int(math.Sqrt(float64(num)))
	for i := upperLimit; i > 1; i-- {
		if num%i == 0 {
			if isPrime(i) {
				fmt.Println(i)
				i = 0
			}
		}
	}
}

func isPrime(num int) bool {

	if num%2 == 0 {
		return false
	} else if num < 3 {
		return false
	}

	//upperLimit := int(math.Sqrt(float64(num)))

	for i := 3; i < num/2; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
