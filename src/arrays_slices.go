package main

import (
	"fmt"
	"strings"
)

func main() {
	//Arrays
	var array [5]int
	array[0] = 11
	array[1] = 25
	array[2] = 9
	fmt.Println(array)

	//Slices
	slice := []int{4, 22, 48, 68, 16, 74, 33, 51}
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[:4])
	fmt.Println(slice[3:5])
	fmt.Println(slice[5:])
	//append new element
	slice = append(slice, 7)
	fmt.Println(slice)
	//append new slice
	newSlice := []int{10, 40, 30}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//slice con range
	sliceStr := []string{"hello", "what", "are", "you", "doing"}
	for i, valor := range sliceStr {
		fmt.Println(i, valor)
	}
	text := "Ama"
	if isPalindrome(text) {
		fmt.Printf("%v it's a palindrome word/phrase", text)
	} else {
		fmt.Printf("%v it's not a palindrome word/phrase", text)
	}

}
func isPalindrome(text string) bool {
	text = strings.ToLower(text)
	var textReverse string
	for j := len(text) - 1; j >= 0; j-- {
		textReverse += string(text[j])
	}
	if text == textReverse {
		return true
	}
	return false
}
