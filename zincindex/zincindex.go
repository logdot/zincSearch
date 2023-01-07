// Package zincindex implements methods for indexing the enron mail database into ZincSearch.
package zincindex

import (
	"bytes"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

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
	a.indexFolder(path)
}

// ClearDB deletes all records from ZincSearch.
// This is a destructive operation and should only be used when absolutely necessary (say, for testing purposes)
func (a *Authentication) ClearDB() {
}

// indexFolder takes a path to the enron database and indexes all of its directories and files recursively.
func (a *Authentication) indexFolder(path string) {
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		mail, err := a.indexFile(path)
		if err != nil {
			return err
		}

		err = a.ingestSingle(mail)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

// indexFile indexes a single file
func (a *Authentication) indexFile(path string) (Mail, error) {
	mail := Mail{}

	file, err := os.Open(path)
	if err != nil {
		return mail, err
	}

	// Defer closing the file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	mail = ParseMailFromFile(file)

	return mail, nil
}

// ingestSingle takes a mail and sends it to be ingested
func (a *Authentication) ingestSingle(mail Mail) error {
	data, err := json.Marshal(mail)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", a.address, bytes.NewReader(data))
	if err != nil {
		log.Println(err)
		return err
	}

	req.SetBasicAuth(a.username, a.password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	defer resp.Body.Close()

	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(string(body))

	return nil
}
