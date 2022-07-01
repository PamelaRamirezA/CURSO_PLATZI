package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	numero := 50
	if numero == 45 {
		fmt.Println("hello world")
	} else {
		fmt.Println("number does not fit with the condition")
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("number ", i)
	}

	name_p := "Esternocleidomastoideo"
	//	fmt.Println(string(name_p[0]))

	for i := 0; i < len(name_p); i++ {
		fmt.Println(string(name_p[i]))
	}

	for name_p != "Pamela" { //similar to while
		fmt.Println("name is different from Pamela")
		name_p = "Pamela"
	}

	for j := len(name_p) - 1; j >= 0; j-- {
		fmt.Println(string(name_p[j]))
	}

	value1 := 1
	value2 := 2

	if value1 == 1 {
		fmt.Println("value is 1")
	} else {
		fmt.Println("value is not 1")
	}
	//and
	if value1 == 1 && value2 == 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//or
	if value1 == 0 || value2 == 2 {
		fmt.Println("true one of them")
	}

	//parsing string to int
	value, err := strconv.Atoi("88")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)

	//switch
	switch mod := 4 % 2; mod {
	case 0:
		fmt.Println("is even")
	default:
		fmt.Println("is odd")
	}
	//without condition
	val := 50
	switch {
	case val > 100:
		fmt.Println("val is greater than 100")
	case val < 0:
		fmt.Println("val is less than 0")
	default:
		fmt.Println("val does not fit with the conditions")
	}

	//defer
	defer fmt.Println("last execution line") //it's used for closing bd connections
	fmt.Println("next line")
	//continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("i is 2")
			continue //the for is stopped here and it starts again on the condition sentence
		}
		if i == 2 {
			fmt.Println("next line after continue")
		}
		if i == 6 {
			fmt.Println("i is 6")
			break //it finishes the for bucle
		}

	}
}
