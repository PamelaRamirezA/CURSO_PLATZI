package mypackage

import "fmt"

//CarPublic is a public structure
type CarPublic struct {
	Brand, Model, Owner string
	Year                int
	Motor               float32
}

//PrintMessage is a public function
func PrintMessage(message string) {
	fmt.Println(message)
}
