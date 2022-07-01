package main

import "fmt"

func main() {
	//number even or odd
	num := 2
	ret := isEven(num)
	if ret {
		fmt.Printf("number %v is even", num)
	} else {
		fmt.Printf("number %v is odd", num)
	}
}

func isEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
