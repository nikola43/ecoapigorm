package utils

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
)

type SendEmailManager struct {
	ToEmail               string
	ToName                string
	FromEmail             string
	FromName              string
	CompanyName           string
	ClinicName           string
	RecoveryToken         string
	InvitationToken       string
	RecoveryPasswordToken string
	StreamingUrl string
	StreamingCode string
	Template string
	Subject string
}

func (i SendEmailManager) SendMail() {
	senderEmail := "ecox.server@stelast.es"
	// senderEmail := GetEnvVariable("FROM_EMAIL")
	// fromEmailPassword := GetEnvVariable("FROM_EMAIL")

	t := template.New(i.Template)
	var err error
	t, err = t.ParseFiles(i.Template)
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", i.ToEmail)
	m.SetHeader("Subject", i.Subject)
	m.SetBody("text/html", result)
	//m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("send.one.com", 465, "ecox.server@stelast.es", "NRSeEKDHK7W6rwDc")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
