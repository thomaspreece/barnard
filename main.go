package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"

	"layeh.com/barnard/uiterm"
	"layeh.com/gumble/gumble"
	_ "layeh.com/gumble/opus"
)

func main() {
	// Command line flags
	server := flag.String("server", "localhost:64738", "the server to connect to")
	username := flag.String("username", "", "the username of the client")
	password := flag.String("password", "", "the password of the server")
	channel := flag.String("channel", "", "the channel to move to after connecting")
	insecure := flag.Bool("insecure", false, "skip server certificate verification")
	certificate := flag.String("certificate", "", "PEM encoded certificate and private key")

	flag.Parse()

	// Initialize
	b := Barnard{
		Config:  gumble.NewConfig(),
		Address: *server,
	}

	b.Config.Username = *username
	b.Config.Password = *password

	b.StartChannel = *channel

	if *insecure {
		b.TLSConfig.InsecureSkipVerify = true
	}
	if *certificate != "" {
		cert, err := tls.LoadX509KeyPair(*certificate, *certificate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		b.TLSConfig.Certificates = append(b.TLSConfig.Certificates, cert)
	}

	b.Ui = uiterm.New(&b)
	b.Ui.Run()
}
