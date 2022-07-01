package main

import (
	mypackage "curso_platzi/src/myPackage"
	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "volkswagen"
	myCar.Year = 2021
	fmt.Println(myCar)

	mypackage.PrintMessage("messageNew")

}
