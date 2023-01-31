package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(5,5))
}

func sum (numA, numB int) (int, error) {
	if numA + numB > 10 {
		return 0, fmt.Errorf("This sum is greater than 10.")
	}
	return numA + numB, nil
}