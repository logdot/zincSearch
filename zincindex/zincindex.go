// Package zincindex implements methods for indexing the enron mail database into ZincSearch.
package zincindex

// Authentication represents a set of credentials for the ZincSearch API.
type Authentication struct {
	address  string
	username string
	password string
}

// Authenticate creates an authentication object from an address, username and password.
func Authenticate(address string, username string, password string) *Authentication {
	authentication := &Authentication{
		address:  address,
		username: username,
		password: password,
	}

	return authentication
}

// IndexDB indexes the enron database at the given path and ingests it into ZincSearch.
// path is expected to point at a directory containing `maildir/`.
func (a *Authentication) IndexDB(path string) {
}

// ClearDB deletes all records from ZincSearch.
// This is a destructive operation and should only be used when absolutely necessary (say, for testing purposes)
func (a *Authentication) ClearDB() {
}
