package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"gopkg.in/gomail.v2"
)

type Info struct {
	Data string
}

func (i Info) SendMailRecovery(htmlTemplate string, email string, subject string) {
	t := template.New(htmlTemplate)
	var err error
	t, err = t.ParseFiles(htmlTemplate)
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "ecodadystest@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	//m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("smtp.gmail.com", 587, "ecodadystest@gmail.com", "#ecodadys1")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
