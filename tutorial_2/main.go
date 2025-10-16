// DIVISION,reminder
// Nil funtion error
// modular function
// If else
// Switch case

package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "Hello, World!"
	printMe(printValue)

	numerator := 11
	denominator := 2
	result, remainder, err := intDivision(numerator, denominator)

	switch {
	case err != nil:
		fmt.Printf("%s", err.Error())
	case remainder == 0:
		fmt.Printf("the result is %v", result)
	default:
		fmt.Printf("the result is %v and reminder is %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("THE DIVISION WAS ZERO")
	case 1, 2:
		fmt.Printf("THE DIVISION WAS close")
	default:
		fmt.Printf("THE DIVISION WAS FAR")
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}
