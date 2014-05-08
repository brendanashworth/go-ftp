package ftp

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

// FTPClient is a representation of a client connected to the FTP server. It contains all data necessary for the entire browsing session.
type FTPClient struct {
	server 			*FTPServer
	conn 			net.Conn // connection instance
	writer 			*bufio.Writer
	scanner 		*bufio.Scanner
	authenticated   bool
	user 			string
	password 		string
	dir				string
	relativedir		string
	transferType 	string
	Mode			int
	dataSocket 		net.Conn // data socket
}

// HandleRequest is used to handle a command sent by the client to the server. It takes a string as a paramater and does not return any data.
func (this *FTPClient) HandleRequest(req string) {
	// get COMMAND, then MESSAGE
	request := strings.SplitAfterN(req, ` `, 2)
	command := strings.Trim(request[0], ` `)

	// did they even send a message?
	if len(request) > 1 {
		message := strings.Trim(request[1], ` `)

		fmt.Println("Command: " + command + ", message: " + message)

		// lets assign the command
		switch command {
		case "USER":
			this.USER(message)
		case "PASS":
			this.PASS(message)
		case "TYPE":
			this.TYPE(message)
		default:
			this.NOTIMP()
		}

	// there was no message
	} else {
		fmt.Println("Command: " + command)

		// handle
		switch command {
		case "PASV":
			this.PASV()
		case "QUIT":
			this.QUIT()
		case "SYST":
			this.SYST()
		case "FEAT":
			this.FEAT()
		case "PWD":
			this.PWD()
		case "LIST":
			this.LIST()
		default:
			this.NOTIMP()
		}
	}
}

// Send a message to the FTP Client.
func (this *FTPClient) SendMessage(code int) {
	message := GetMessages()[code]
	completeMsg := strconv.Itoa(code) + " " + message

	this.Write(completeMsg)
	fmt.Println(completeMsg)
}

// Send a message to the FTP Client, with an injectable.
func (this *FTPClient) SendMessageWithInjectable(code int, injectable string) {
	message := GetMessages()[code]
	completeMsg := strconv.Itoa(code) + " " + message
	completeMsg = strings.Replace(completeMsg, "%s", injectable, -1)

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

// Write a string to the client's data socket.
func (this *FTPClient) WriteDataSocket(message string) {
	writer := bufio.NewWriter(this.dataSocket)
	_, err := writer.WriteString(message + "\n")
	if err != nil {
		fmt.Println("Error occurred writing to data socket connection: " + err.Error())
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error occurred flushing data socket stream: " + err.Error())
	}
}

// Closes the data socket.
func (this *FTPClient) CloseDataSocket() {
	this.dataSocket.Close()
}

// Closes the client.
func (this *FTPClient) Close() {
	this.conn.Close()
}