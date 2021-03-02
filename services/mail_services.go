package services

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
)

func SendMailRecovery(umail string, token string){
	t := template.New("recovery_password.html") // ¿Porque esta doble recovery_password.html

	var err error
	t, err = t.ParseFiles("recovery_password.html")
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, umail);
	err != nil {
		log.Println(err)
	}

	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "mimatrona.soporte@gmail.com")
	m.SetHeader("To", umail)
	m.SetHeader("Subject", "Cambiar contraseña")
	m.SetBody("text/html", result)
	//m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("smtp.gmail.com", 587, "mimatrona.soporte@gmail.com", "#ecodadys1")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
