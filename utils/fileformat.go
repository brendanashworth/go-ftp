package utils

import (
	"os"
	"strconv"
)

// Parses a file into the correct FTP-defined string.
func ParseFile(file os.FileInfo, parseType string) string {
	fileName := file.Name()
	bytes := strconv.FormatInt(file.Size(), 10)

	switch parseType {
	// ASCII data format type
	case "A":
		var host string
		// permissions
		if file.IsDir() {
			host = host + "drwxr-xr-x"
		} else {
			host = host + "-rw-r--r--"
		}

		// space separator
		host = host + " "

		// id, owner, group
		host = host + "1 owner group"

		// byte size
		byteSize := 13 - len(bytes)
		for byteSize > 0 {
			byteSize = byteSize - 1
			host = host + " "
		}
		host = host + bytes

		// temporary date fix
		host = host + " Jul 27 01:45 "

		// filename
		host = host + fileName

		return host
	default:
		return "That data format type is not yet supported."
	}
}