package main

import "fmt"

func say(text string) {
	fmt.Println(text)
}

func message(text string, c chan string) {
	c <- text
}
func main() {
	chan1 := make(chan string, 3)
	chan1 <- "Message 1"
	chan1 <- "Message 2"

	fmt.Println(len(chan1), cap(chan1))

	close(chan1) //the channel is closed in order to not received any other data

	//range
	for message := range chan1 {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go message("Message1", email1)
	go message("Message2", email2)

	for i := 0; i < 2; i++ { //it's important to know the number of the channels to establish the max value in for routine
		select {
		case rpta1 := <-email1:
			fmt.Println("Email1 ", rpta1)
		case rpta2 := <-email2:
			fmt.Println("Email2 ", rpta2)
		}

	}
}
