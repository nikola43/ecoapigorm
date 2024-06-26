package utils

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"crypto/tls"
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

	d := gomail.NewDialer("cph-ecox4d.customprofessionalhosting.com", 587, "enviosapp@ecox4d.com", "Ff86f?n0")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
