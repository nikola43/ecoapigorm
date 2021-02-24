package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	gomail "gopkg.in/gomail.v2"
)

type Info struct {
	Data string
}

func (i Info) SendMailRecovery(umail string) {
	t := template.New("recovery_password.html")

	var err error
	t, err = t.ParseFiles("recovery_password.html")
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
	m.SetHeader("To", umail)
	m.SetHeader("Subject", "Cambiar contraseña")
	m.SetBody("text/html", result)
	//m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("smtp.gmail.com", 587, "ecodadystest@gmail.com", "#ecodadys1")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

func (i Info) SendStreamingMail(umail string) {
	t := template.New("new_streaming.html")

	var err error
	t, err = t.ParseFiles("new_streaming.html")
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
	m.SetHeader("To", umail)
	m.SetHeader("Subject", "Nuevo streaming disponible")
	m.SetBody("text/html", result)
	//m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("smtp.gmail.com", 587, "ecodadystest@gmail.com", "#ecodadys1")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

func (i Info) SendWelcomeMail(umail string) {
	t := template.New("welcome.html")

	var err error
	t, err = t.ParseFiles("welcome.html")
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
	m.SetHeader("To", umail)
	m.SetHeader("Subject", "Bienvenid@ a eCodadys")
	m.SetBody("text/html", result)
	//m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("smtp.gmail.com", 587, "ecodadystest@gmail.com", "#ecodadys1")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

func (i Info) SendPasswordRecoveryConfirmation(umail string) {
	t := template.New("password_recovery_confirmation.html")

	var err error
	t, err = t.ParseFiles("password_recovery_confirmation.html")
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
	m.SetHeader("To", umail)
	m.SetHeader("Subject", "Se ha cambiado su contraseña")
	m.SetBody("text/html", result)
	//m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("smtp.gmail.com", 587, "ecodadystest@gmail.com", "#ecodadys1")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
