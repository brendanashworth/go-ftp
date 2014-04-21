package example

import (
	"fmt"
	"github.com/boboman13/go-ftp"
)

func main() {
	server := &FTPServer{
		host: "0.0.0.0",
		port: "21",
	}

	err := server.Start()
	if err != nil {
		fmt.Println("Error occurred starting FTP server: " + err.Error())
	}
}