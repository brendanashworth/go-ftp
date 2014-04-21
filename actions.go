package main

import (
	"fmt"
	"strings"
)

// Client QUIT
func (this *FTPClient) QUIT() {
	this.SendMessage(221)
	this.Close()
}

// Client USER
func (this *FTPClient) USER(message string) {
	// grab username
	username := strings.Trim(message, " ")
	this.user = username

	this.SendMessage(331)
}

// Client PASS
func (this *FTPClient) PASS(message string) {
	// grab password
	password := message
	this.password = password

	this.SendMessage(230)
}

// Client SYST
func (this *FTPClient) SYST() {
	this.SendMessage(215)
}

func (this *FTPClient) CWD(message string) {
	fmt.Println("Changed working directory to: " + message)
}

