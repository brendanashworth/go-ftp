package main

import (
	"fmt"
)

// Client QUIT
func (this *FTPClient) QUIT() {
	this.SendMessage(221)
	this.Close()
	fmt.Println("Closed connection")
}

func (this *FTPClient) CWD(message string) {
	fmt.Println("Changed working directory to: " + message)
}

func (this *FTPClient) PASS(message string) {
	fmt.Println("Set password to: " + message)
}