package main

import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

//SMS
type SMSNotification struct {
	destinatary string
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notificacion via SMS")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

//Email
type EmailNotification struct {
	destinataryEmail string
}

func (EmailNotification) SendNotification() {
	fmt.Println("sending notification via Email")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func main() {
	sms := SMSNotification{}
	email := EmailNotification{}
	fmt.Println("SMS method Processing")
	sms.SendNotification()
	fmt.Println("Method ", sms.GetSender().GetSenderMethod())
	fmt.Println("Channel ", sms.GetSender().GetSenderChannel())

	fmt.Println("\nEmail method Processing")
	email.SendNotification()
	fmt.Println("Method ", email.GetSender().GetSenderMethod())
	fmt.Println("Channel ", email.GetSender().GetSenderChannel())
}
