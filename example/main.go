package main

import (
	"fmt"
	"github.com/boboman13/go-ftp/ftp"
)

// Starts server
func main() {
	server := &ftp.FTPServer{
		Host: "0.0.0.0",
		Port: 21,
		Config: new(ftp.AuthenticationConfig),
	}

	// Configures authentication; WARNING: do not use this code, it is insecure
	server.Config.ConfigAuthentication(func(user string, password string) (authenticated bool, dir string) {
		fmt.Println("Logged in " + user + " w/ pass " + password)
		return true, "/home/" + user + "/ftp"
		})

	// Starts server
	err := server.Start()
	if err != nil {
		fmt.Println("Error occurred starting FTP server: " + err.Error())
	}
}