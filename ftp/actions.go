package ftp

import (
	"net"
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
	"math/rand"
	"github.com/boboman13/go-ftp/utils"
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

// Client PASV
func (this *FTPClient) PASV() {
	this.Mode = PASSIVE_MODE

	port := rand.Intn(65535)
	listener, err := net.Listen("tcp", this.server.Host + ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error launching PASSIVE port: " + err.Error())
		this.Close()
	}

	// send message to client
	host := strings.Replace(this.server.Host, ".", ",", -1)
	extraport := port % 256
	primaryport := (port - extraport) / 256

	host = host + "," + strconv.Itoa(primaryport) + "," + strconv.Itoa(extraport)
	this.SendMessageWithInjectable(227, host)

	// get our one client
	conn, err := listener.Accept()
	this.dataSocket = conn
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

// Client TYPE, basically changes format to display files
func (this *FTPClient) TYPE(message string) {
	this.transferType = message
	this.SendMessage(200)
}

// Client LIST, basically a directory listing
func (this *FTPClient) LIST() {
	// make sure this client is authenticated!
	if !this.authenticated {
		this.SendMessage(503)
		return
	}

	this.SendMessage(150)

	// gets files
	files, err := ioutil.ReadDir(this.dir + this.relativedir)
	if err != nil {
		fmt.Println("Error while reading directory ("+ this.dir + this.relativedir +"): " + err.Error())

		this.SendMessage(226)
		return
	}

	fmt.Println(this.dir + this.relativedir)

	for _, file := range(files) {
		format := utils.ParseFile(file, this.transferType)
		this.WriteDataSocket(format)
	}
	this.CloseDataSocket()

	this.SendMessage(226)
}

// Client had sent a command not implemented.
func (this *FTPClient) NOTIMP() {
	this.SendMessage(502)
}