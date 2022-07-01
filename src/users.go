package main

import "fmt"

func main() {
	user1 := "Pamela"
	pass1 := "Pamela123"
	if isUserValid(user1, pass1) {
		fmt.Println("Access approved")
	} else {
		fmt.Println("Access denied")

	}
}

func isUserValid(userName string, password string) bool {
	if userName == "Pamela" && password == "Pra123" {
		return true
	} else {
		return false
	}

}
