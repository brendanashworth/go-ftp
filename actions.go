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
	// check if there is a username
	if this.user == "" {
		this.SendMessage(503)
		return
	}

	// check if not authentication
	if this.authenticated {
		this.SendMessage(503)
		return
	}

	// grab password
	password := message
	this.password = password

	// authenticate
	authenticated, dir := this.server.config.authenticationCheck(this.user, this.password)
	this.authenticated = authenticated
	this.dir = dir
	
	if this.authenticated {
		this.SendMessage(230)
	} else {
		this.SendMessage(230)
	}
}

// Client SYST
func (this *FTPClient) SYST() {
	this.SendMessage(215)
}

// Client FEAT
func (this *FTPClient) FEAT() {
	this.SendMessage(502)
}

func (this *FTPClient) CWD(message string) {
	fmt.Println("Changed working directory to: " + message)
}

// Client, command not implemented.
func (this *FTPClient) NOTIMP() {
	this.SendMessage(502)
}