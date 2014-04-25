package ftp

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
)

// A representation of the FTPServer. It contains a net.listener instance, a host, a port, and a config (see AuthenticationConfig).
type FTPServer struct {
	listener net.Listener
	Host 	 string
	Port 	 int
	Config	 *AuthenticationConfig
}

// Instantiates and returns a new FTPServer instance. This function does not start the server and does not block the main thread.
func CreateServer(host string, port int) (server *FTPServer) {
	server = &FTPServer{
		Host: host,
		Port: port,
		Config: new(AuthenticationConfig),
	}

	return server
}

// Starts an FTP server. Must already have a FTPServer instance. This method is blocking and should be used within a goroutine if asynchronous ability
// is expected.
func (this *FTPServer) Start() (err error) {
	listener, err := net.Listen("tcp", this.Host + ":" + strconv.Itoa(this.Port))
	if err != nil {
		return err
	}

	this.listener = listener

	fmt.Println("FTP Server started on " + this.Host + ":" + strconv.Itoa(this.Port) + ".")

	// lets go into a for loop for clients
	for {
		conn, err := this.listener.Accept()
		if err != nil {
			fmt.Println("Error occurred accepting FTP client: " + err.Error())
			continue
		}

		client := &FTPClient{
			conn: conn,
			authenticated: false,
			server: this,
		}

		go this.HandleClient(client)
	}

	return nil
}

// This function is used for handling clients, and should be called inside a goroutine if asynchronous ability is expected. This function takes in one
// paramater, an FTPClient instance.
func (this *FTPServer) HandleClient(client *FTPClient) {
	fmt.Println("Now handling client: " + client.conn.RemoteAddr().String())

	// get us a scanner
	client.scanner = bufio.NewScanner(client.conn)
	client.writer = bufio.NewWriter(client.conn)

	// send welcome message, then wait for text
	client.SendMessage(220)

	for client.scanner.Scan() {
		cmd := client.scanner.Text()
		client.HandleRequest(cmd)
	}
}