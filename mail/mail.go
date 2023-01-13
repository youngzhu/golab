package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

const (
	sender    = "1"
	senderPwd = "1"
	reveiver  = "1"
)

func main() {
	m := gomail.NewMessage()

	m.SetHeader("From", sender)
	m.SetHeader("To", reveiver)

	m.SetHeader("Subject", "Gomail test")

	m.SetBody("text/plain", "This is Gomail test")

	d := gomail.NewDialer("smtp.163.com", 465, sender, senderPwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("OK")

}
