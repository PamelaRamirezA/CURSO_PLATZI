package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func newPc(ram int, disk int, brand string) *pc {
	return &pc{
		ram:   ram,
		disk:  disk,
		brand: brand,
	}
}
func (value pc) ping() {
	fmt.Println(value.brand, "Pong")
}

func (value *pc) duplicateRam() {
	value.ram *= 2
}

func (value pc) String() string { //String is a method that is part of struct which can be overwritten
	return fmt.Sprintf("My brand is %s, I have  %d RAM, and  %dGB of hard disk", value.brand, value.ram, value.disk)
}

func main() {
	a := 50
	b := &a

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("*b", *b)

	//mypc := pc{ram: 16, disk: 200, brand: "ASUS"}
	mypc := newPc(8, 100, "ASUS2")
	fmt.Println(mypc)

	mypc.ping()
	fmt.Println(mypc)
	mypc.duplicateRam()
	fmt.Println(mypc)
	mypc.duplicateRam()
	fmt.Println(mypc)
}
