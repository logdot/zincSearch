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
	index    string
	username string
	password string
}

// Authenticate creates an authentication object from an address, username and password.
func Authenticate(address string, index string, username string, password string) *Authentication {
	authentication := &Authentication{
		address:  address,
		index:    index,
		username: username,
		password: password,
	}

	return authentication
}

// IndexDB indexes the enron database at the given path and ingests it into ZincSearch.
// path is expected to point at a directory containing `maildir/`.
func (a *Authentication) IndexDB(path string) {
	log.Println("Starting indexing")

	mails := indexFolder(path)

	err := a.ingestBulk(mails, 10000)

	if err != nil {
		log.Println(err)
	}

	//for _, mail := range mails {
	//	err := a.ingestSingle(mail)
	//
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}
}

// ClearDB deletes all records from ZincSearch.
// This is a destructive operation and should only be used when absolutely necessary (say, for testing purposes)
func (a *Authentication) ClearDB() {
}

// indexFolder takes a path to the enron database and indexes all of its directories and files recursively.
func indexFolder(path string) []Mail {
	var mails []Mail

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		mail, err := indexFile(path)
		if err != nil {
			return err
		}

		mails = append(mails, mail)

		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return mails
}

// indexFile indexes a single file
func indexFile(path string) (Mail, error) {
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

	requestAddr := a.address + "api/" + a.index + "/_doc"
	return a.sendRequest(requestAddr, data)
}

func (a *Authentication) ingestBulk(mails []Mail, chunking int) error {
	index := struct {
		InnerIndex struct {
			Index string `json:"_index"`
		} `json:"index"`
	}{
		InnerIndex: struct {
			Index string `json:"_index"`
		}{
			Index: a.index,
		},
	}

	indexData, err := json.Marshal(index)
	if err != nil {
		return err
	}
	// ZincSearch expects a newline delimited json file
	indexData = append(indexData, '\n')

	requestBody := []byte("")
	for i, mail := range mails {
		// For some reason if we want to index data we need to specify the index for every single item, instead of only
		// once at the beginning of the request.
		requestBody = append(requestBody, indexData...)

		data, err := json.Marshal(mail)
		if err != nil {
			return err
		}

		// ZincSearch expects a newline delimited json file
		data = append(data, '\n')

		requestBody = append(requestBody, data...)

		if i%chunking == 0 {
			requestAddr := a.address + "api/_bulk"
			err = a.sendRequest(requestAddr, requestBody)
			if err != nil {
				return err
			}
			requestBody = []byte("")
		}
	}

	requestAddr := a.address + "api/_bulk"
	return a.sendRequest(requestAddr, requestBody)
}

func (a *Authentication) sendRequest(url string, body []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
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
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(string(body))
	return nil
}
