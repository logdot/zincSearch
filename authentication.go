package zincSearch

// Authentication represents a set of credentials for the ZincSearch API.
type Authentication struct {
	address  string
	index    string
	username string
	password string
}

// Authenticate creates an authentication object from an address, index, username and password.
func Authenticate(address string, index string, username string, password string) *Authentication {
	authentication := &Authentication{
		address:  address,
		index:    index,
		username: username,
		password: password,
	}

	return authentication
}
