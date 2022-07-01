package main

import (
	"fmt"
	"os"

	"github.com/wneessen/go-mail"
)

func main() {
	mail1 := mail.NewMsg()

	if error1 := mail1.FromFormat("Pamela Ramirez", "sramirez.1506@gmail.com"); error1 != nil {
		fmt.Println("failed setting the FROM address", error1)
		os.Exit(1)
	}
	if error1 := mail1.To("Test email", "sramirez.1506@gmail.com"); error1 != nil {
		fmt.Println("failed setting the TO address", error1)
		os.Exit(1)
	}
	mail1.Subject("Subject from the email")
	mail1.SetDate()      //setea la fecha actual en el encabezado
	mail1.SetMessageID() //setea un codigo unico de mensaje
	mail1.SetBulk()      // indica el tamanio del mail
	mail1.SetImportance(mail.ImportanceHigh)
	mail1.SetBodyString(mail.TypeTextPlain, "Body from the email")

}
