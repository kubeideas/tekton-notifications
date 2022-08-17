package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"notifications/email"
	"notifications/errors"
	"notifications/slack"
	"os"
	"strings"
)

//go:embed usage/usage.txt
var usageFile string

func printUsage() {

	fmt.Println(usageFile)

	os.Exit(0)
}

func main() {

	usage := flag.Bool("usage", false, "Print usage instructions")
	flag.Parse()

	if *usage {
		printUsage()
	}

	var sendType string
	var present bool
	if sendType, present = os.LookupEnv("SEND_TYPE"); !present {
		log.Fatalln(&errors.EnvVarNotDefined{Name: "SEND_TYPE"})
	}

	switch strings.ToUpper(sendType) {
	case "SMTP":
		var si email.SmtpInfo
		err := si.SendEmailNotification()
		if err != nil {
			log.Fatalln(err)
		}
	case "SLACK":
		var ji slack.JsonInfo
		err := ji.SendJsonNotification()
		if err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatal(&errors.NotificationNotImplemented{Name: sendType})
	}

}
