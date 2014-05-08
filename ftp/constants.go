package ftp

const (
	// RCV_BUFFER_LENGTH is the size of the buffer used in the FTP server.
	RCV_BUFFER_LENGTH = 1024
	// ACTIVE_MODE is the representation of FTP active mode.
	ACTIVE_MODE = 11
	// PASSIVE_MODE is the representation of FTP passive mode.
	PASSIVE_MODE = 12
)

// GetMessages() is a simple way to get all the messages used commonly by the server in responses to response quickly and easily.
func GetMessages() (messages map[int]string) {
	messages = map[int]string{
		150: "Directory listing incoming.",
		200: "Command successfull.",
		215: "Test unix system.",
		220: "Hello, this is Go-FTP server.",
		221: "Goodbye.",
		226: "Action completed.",
		227: "Entering Passive Mode (%s)",
		230: "Logged in.",
		257: "\"%s\" is current directory.",
		331: "Password required for access to account.",
		502: "Command not implemented.",
		503: "Bad sequence of commands.",
	}

	return messages
}