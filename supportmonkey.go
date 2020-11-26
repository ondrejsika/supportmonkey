// supportmonkey
//
// Author: Ondrej Sika <ondrej@ondrejsika.com>
// Source: https://github.com/ondrejsika/gosendmail
//

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	gosendmail "github.com/ondrejsika/gosendmail/lib"
)

func main() {
	from := flag.String("from", "", "")
	to := flag.String("to", "", "")
	smtpHost := flag.String("smtp-host", "", "")
	smtpPort := flag.String("smtp-port", "587", "optional")
	password := flag.String("password", "", "")
	debug := flag.Bool("debug", false, "optional")
	subject := flag.String("subject", "", "optional")
	message := flag.String("message", "", "optional")

	flag.Parse()

	if *from == "" {
		fmt.Fprintf(os.Stderr, "-from is not defined\n")
		os.Exit(1)
	}
	if *to == "" {
		fmt.Fprintf(os.Stderr, "-to is not defined\n")
		os.Exit(1)
	}
	if *smtpHost == "" {
		fmt.Fprintf(os.Stderr, "-smtp-host is not defined\n")
		os.Exit(1)
	}
	if *smtpPort == "" {
		fmt.Fprintf(os.Stderr, "-smtp-port is not defined\n")
		os.Exit(1)
	}
	if *password == "" {
		fmt.Fprintf(os.Stderr, "-password is not defined\n")
		os.Exit(1)
	}

	for {
		// Random wait
		min := 20
		max := 80
		randomMultiplier := time.Duration(rand.Intn((max - min) + min))
		var sleepTime time.Duration
		sleepTime = time.Hour
		if *debug {
			sleepTime = randomMultiplier * time.Second
		} else {
			sleepTime = randomMultiplier * time.Hour
		}
		fmt.Println("Send email in " + sleepTime.String())
		time.Sleep(sleepTime)

		// Send email
		emailSubject := "[supportmonkey] New incident"
		if *subject != "" {
			emailSubject = "[supportmonkey] " + *subject
		}
		emailMessage := "This is an example incident from support monkey. --supportmonkey"
		if *message != "" {
			emailMessage = *message + " --supportmonkey"
		}
		err := gosendmail.GoSendMail(*smtpHost, *smtpPort, *from, *password, *to, emailSubject, emailMessage)
		if err != nil {
			panic(err)
		}
	}
}
