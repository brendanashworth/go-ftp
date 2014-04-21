package main

import (
	"fmt"
	"net"
	"strconv"
)

// FTP struct
type FTPServer struct {
	listener net.Listener // connection instance
	host string
	port int
}

// FTP client struct
type FTPClient struct {
	conn net.Conn // connection instance
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

	for {
		buf := make([]byte, RCV_BUFFER_LENGTH) // create a buffer

		bytes, err := client.conn.Read(buf)
		if err != nil {
			fmt.Println("Error occurred accepting FTP message: " + err.Error())
			return
		}

		fmt.Println("Received " + strconv.Itoa(bytes) + " bytes of data, equal to " + string(buf))

		// handle the request
		this.HandleRequest(string(buf), client)
	}
}

// This function handles an FTP request.
func (this *FTPServer) HandleRequest(req string, client *FTPClient) {

}

// Main method
