package ftp

// Auth function. Takes in first string as USER, second string as PASSWORD. Returns true / false based on whether or not they should
//   be authenticated. If it returns true, also returns a string where they are starting to browse.
type authCheck func(string, string) (bool, string)

// Authentication struct, used for configuring the authentication server.
type AuthenticationConfig struct {
	authenticationCheck authCheck
}

// Handles authentication for the FTP server.
func (this *AuthenticationConfig) ConfigAuthentication(function authCheck) {
	this.authenticationCheck = function
}