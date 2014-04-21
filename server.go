package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// FTP struct
type FTPServer struct {
	listener net.Listener // connection instance
	host string
	port int
}

func main() {
	server := &FTPServer{
		host: "0.0.0.0",
		port: 21,
	}

	err := server.Start()
	if err != nil {
		fmt.Println("Error occurred starting FTP server: " + err.Error())
	}
}

// Starts an FTP server. Must already have a FTPServer instance. This method is blocking.
func (this *FTPServer) Start() (err error) {
	// listen
	listener, err := net.Listen("tcp", this.host + ":" + strconv.Itoa(this.port))
	if err != nil {
		return err
	}

	this.listener = listener

	fmt.Println("FTP Server started on " + this.host + ":" + strconv.Itoa(this.port) + ".")

	// lets go into a for loop for clients
	for {
		conn, err := this.listener.Accept()
		if err != nil {
			fmt.Println("Error occurred accepting FTP client: " + err.Error())
			continue
		}

		client := &FTPClient{
			conn: conn,
		}

		go this.HandleClient(client)
	}

	return nil
}

// This function is used for handling clients, and should be called inside a go routine.
func (this *FTPServer) HandleClient(client *FTPClient) {
	fmt.Println("Now handling client: " + client.conn.RemoteAddr().String())

	// send welcome message
	client.SendMessage(220)

	for {
		buf := make([]byte, RCV_BUFFER_LENGTH) // create a buffer

		_, err := client.conn.Read(buf)
		if err != nil {
			fmt.Println("Error occurred accepting FTP message: " + err.Error())
			return
		}

		// handle the request
		this.HandleRequest(string(buf), client)
	}
}

// This function handles an FTP request.
func (this *FTPServer) HandleRequest(req string, client *FTPClient) {
	// get COMMAND, then MESSAGE
	request := strings.SplitAfterN(req, ` `, 2)

	command := strings.Trim(request[0], ` `)

	// did they even send a message?
	if len(request) > 1 {
		message := request[1]

		fmt.Println("Command: " + command + ", message: " + message)

		// lets assign the command
		switch command {
		case "CWD":
			client.CWD(message)
		case "PASS":
			client.PASS(message)
		case "QUIT":
			client.QUIT()
		}

	// there was no message
	} else {
		fmt.Println("Command: " + command)

		// handle
		switch command {
		case "QUIT":
			client.QUIT()
		}
	}
}