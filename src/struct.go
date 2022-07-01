package main

import (
	"fmt"
)

type car struct {
	brand, model, owner string
	year                int
	motor               float32
}

func (value car) getModel() string {
	return value.model
}

func (value car) getMotor() float32 {
	return value.motor
}

func main() {

	//first way of instance
	myCar := car{brand: "Hyundai", year: 2020, model: "Accent", owner: "Pamela", motor: 1.6}
	fmt.Println(myCar)

	//second way of instance
	var myCar2 car
	myCar2.brand = "Volkswagen"
	myCar2.year = 2006

	fmt.Println(myCar2)

	fmt.Println(myCar.getMotor())

	fmt.Println(myCar.getModel())
}
