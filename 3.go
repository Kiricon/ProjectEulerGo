package main

import "fmt"

func main() {
	num := 13195

	for i := num; i > 1; i-- {
		if num%i == 0 {
			if isPrime(i) {
				fmt.Println(i)
			}
		}
	}
}

func isPrime(num int) bool {

	if num%2 == 0 {
		return false
	}

	for i := 3; i < num/2; num += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
