package file

import (
	"os"
	"strings"
	"strconv"
	"github.com/boboman13/go-ftp/ftp"
)

// Parses a file into the correct FTP-defined string.
func ParseFile(file os.FileInfo, parseType string) string {
	fileName := file.Name()
	bytes := strconv.FormatUint(file.Size(), 10)

	switch parseType {
	// ASCII data format type
	case "A":

		//      drwx------   3 slacker    users         104 Jul 27 01:45 public_html"   <- example
		return "drwx------   3 slacker    users         " + bytes + " Jul 27 01:45 " + fileName
	}
}