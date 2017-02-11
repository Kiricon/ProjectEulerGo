package main

import "fmt"

func main() {

	currNum := 1
	lastNum := 1
	total := 0

	for currNum < 4000000 {
		if currNum%2 == 0 {
			total += currNum
		}

		temp := currNum
		currNum += lastNum
		lastNum = temp
	}

	fmt.Println(total)

}
