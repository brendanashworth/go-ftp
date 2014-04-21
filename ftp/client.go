package ftp

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

// FTP client struct
type FTPClient struct {
	server 			*FTPServer
	conn 			net.Conn // connection instance
	writer 			*bufio.Writer
	scanner 		*bufio.Scanner
	authenticated   bool
	user 			string
	password 		string
	dir				string
}

func GetMessages() (messages map[int]string) {
	messages = map[int]string{
		200: "PORT command successfull.",
		215: "Test unix system.",
		220: "Hello, this is Go-FTP server.",
		221: "Goodbye.",
		226: "Action completed.",
		230: "Logged in.",
		257: "\"%s\" is current directory.",
		331: "Password required for access to account.",
		502: "Command not implemented.",
		503: "Bad sequence of commands.",
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

// Send a message to the FTP Client, with an injectable.
func (this *FTPClient) SendMessage(code int, injectable string) {
	message := GetMessages()[code]
	completeMsg := strconv.Itoa(code) + " " + message
	strings.Replace(completeMsg, "%s", injectable, -1)

	this.Write(completeMsg)
	fmt.Println(completeMsg)
}

// Write a string to the client.
func (this *FTPClient) Write(message string) {
	_, err := this.writer.WriteString(message + "\n")
	if err != nil {
		fmt.Println("Error occurred writing to connection: " + err.Error())
		return
	}
	err = this.writer.Flush()
	if err != nil {
		fmt.Println("Error occurred flushing data stream: " + err.Error())
	}
}

// Closes the client.
func (this *FTPClient) Close() {
	this.conn.Close()
}