package email

import (
	"bytes"
	"log"
	"net/smtp"
	"notifications/errors"
	"os"
	"strings"
)

type SmtpInfo struct {
	Server     string
	Port       string
	Username   string
	Password   string
	From       string
	Recipients []string
	Msg        []byte
	Subject    string
}

// Load template information fr
func (si *SmtpInfo) setSmtpInfo() error {
	var present bool

	if si.Server, present = os.LookupEnv("SMTP_SERVER"); !present {
		return &errors.EnvVarNotDefined{Name: "SMTP_SERVER"}
	}

	if si.Port, present = os.LookupEnv("SMTP_PORT"); !present {
		return &errors.EnvVarNotDefined{Name: "SMTP_PORT"}
	}

	if si.Username, present = os.LookupEnv("SMTP_USERNAME"); !present {
		return &errors.EnvVarNotDefined{Name: "SMTP_USERNAME"}
	}

	if si.Password, present = os.LookupEnv("SMTP_PASSWORD"); !present {
		return &errors.EnvVarNotDefined{Name: "SMTP_PASSWORD"}
	}

	if si.Subject, present = os.LookupEnv("SMTP_SUBJECT"); !present {
		return &errors.EnvVarNotDefined{Name: "SMTP_SUBJECT"}
	}

	if si.From, present = os.LookupEnv("SMTP_FROM"); !present {
		return &errors.EnvVarNotDefined{Name: "SMTP_FROM"}
	}

	var recipients string
	if recipients, present = os.LookupEnv("SMTP_RECIPIENTS"); !present {
		return &errors.EnvVarNotDefined{Name: "SMTP_RECIPIENTS"}
	} else {
		si.Recipients = strings.Split(recipients, ",")
	}

	return nil
}

// Send mail notification
func (si *SmtpInfo) SendEmailNotification() error {

	var tplBuffer bytes.Buffer

	//parse email html
	var et EmailHtmlTemplate

	err := et.ParseEmailTemplate(&tplBuffer)
	if err != nil {
		log.Fatalln(err)
	}

	// set smtp info
	err = si.setSmtpInfo()
	if err != nil {
		return err
	}

	si.Msg = []byte("Subject: " + si.Subject + "\r\n" + "Content-Type: text/html;charset=utf-8\r\n" +
		tplBuffer.String() + "\r\n")

	auth := smtp.PlainAuth("", si.Username, si.Password, si.Server)

	err = smtp.SendMail(si.Server+":"+si.Port, auth, si.From, si.Recipients, si.Msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("email delivered to [ %s ]", strings.Join(si.Recipients, ","))

	return nil
}
