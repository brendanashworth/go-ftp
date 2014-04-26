package ftp

import (
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
	authenticated, dir := this.server.Config.authenticationCheck(this.user, this.password)
	this.authenticated = authenticated
	this.dir = dir
	this.relativedir = "/"
	
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

// Client PWD
func (this *FTPClient) PWD() {
	this.SendMessageWithInjectable(257, this.relativedir)
}

// Client TYPE
func (this *FTPClient) TYPE(message string) {
	this.transferType = message
	this.SendMessage(200)
}

func (this *FTPClient) PASV() {
	this.SendMessage(150)
	this.Write("drwx------   3 slacker    users         104 Jul 27 01:45 public_html")
	this.SendMessage(226)
}

// Client had sent a command not implemented.
func (this *FTPClient) NOTIMP() {
	this.SendMessage(502)
}