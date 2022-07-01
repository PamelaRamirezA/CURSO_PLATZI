package main

import "fmt"

func main() {
	//maps
	m := make(map[string]float64)

	m["potatoes"] = 3.25
	m["onions"] = 1.89
	m["oranges"] = 0.85

	fmt.Println(m)
	//another way to create a map
	dic := map[string]int{
		"Jose":   8,
		"Pamela": 32,
		"Luis":   28,
	}
	fmt.Println(dic)

	//Looking for a key in a map
	value, ok := dic["Pamela"]
	fmt.Println(value, ok)
}
