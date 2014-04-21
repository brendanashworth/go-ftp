package main

import (
	"net"
	"fmt"
	"bufio"
	"strconv"
)

// FTP client struct
type FTPClient struct {
	conn 			net.Conn // connection instance
	writer 			*bufio.Writer
	scanner 		*bufio.Scanner
	authenticated   bool
	user 			string
	password 		string
}

func GetMessages() (messages map[int]string) {
	messages = map[int]string{
		200: "PORT command successfull.",
		215: "Test unix system.",
		220: "Hello, this is Go-FTP server.",
		221: "Goodbye.",
		226: "Action completed.",
		230: "Logged in.",
		331: "Password required for access to account.",
	}

	return messages
}

// Send a message to the FTP Client.
func (this *FTPClient) SendMessage(code int) {
	message := GetMessages()[code]
	completeMsg := strconv.Itoa(code) + " " + message

	this.Write(completeMsg)
	fmt.Println(completeMsg)
}

// Write a string to the client.
func (this *FTPClient) Write(message string) {
	_, err := this.writer.WriteString(message + "\n")
	if err != nil {
		fmt.Println("Error occurred writing to connection: " + err.Error())
	}
}

// Closes the client.
func (this *FTPClient) Close() {
	this.conn.Close()
}