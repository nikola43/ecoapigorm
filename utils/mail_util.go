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
}

func (i SendEmailManager) SendMail(htmlTemplate string, subject string) {
	senderEmail := "babyandme@stelast.com"
	// senderEmail := GetEnvVariable("FROM_EMAIL")
	// fromEmailPassword := GetEnvVariable("FROM_EMAIL")

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
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", i.ToEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	//m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("ssl0.ovh.net", 465, "mimatrona@stelast.com", "T<NaRMT7}skS4jnQ")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
