package ftp

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
	"strings"
)

// FTP struct
type FTPServer struct {
	listener net.Listener // connection instance
	host 	 string
	port 	 int
	config	 *AuthenticationConfig
}

func main() {
	server := &FTPServer{
		host: "0.0.0.0",
		port: 21,
		config: new(AuthenticationConfig),
	}

	server.config.ConfigAuthentication(func(user string, password string) (authenticated bool, dir string) {
		fmt.Println("Logged in " + user + " w/ pass " + password)
		return true, "/root/go/ftp/test"
		})

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
			authenticated: false,
			server: this,
		}

		go this.HandleClient(client)
	}

	return nil
}

// This function is used for handling clients, and should be called inside a go routine.
func (this *FTPServer) HandleClient(client *FTPClient) {
	fmt.Println("Now handling client: " + client.conn.RemoteAddr().String())

	// get us a scanner
	client.scanner = bufio.NewScanner(client.conn)
	client.writer = bufio.NewWriter(client.conn)

	// send welcome message, then wait for text
	client.SendMessage(220)

	for client.scanner.Scan() {
		cmd := client.scanner.Text()
		this.HandleRequest(cmd, client)
	}
}

// This function handles an FTP request.
func (this *FTPServer) HandleRequest(req string, client *FTPClient) {
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
			client.USER(message)
		case "PASS":
			client.PASS(message)
		default:
			client.NOTIMP()
		}

	// there was no message
	} else {
		fmt.Println("Command: " + command)

		// handle
		switch command {
		case "QUIT":
			client.QUIT()
		case "SYST":
			client.SYST()
		case "FEAT":
			client.FEAT()
		default:
			client.NOTIMP()
		}
	}
}