package zincSearch

import (
	"bytes"
	"encoding/json"
	"github.com/logdot/zincSearch/internal/zincindex"
	"io"
	"log"
	"net/http"
)

// IngestSingle takes a mail and sends it to be ingested
func (a *Authentication) IngestSingle(mail zincindex.Mail) error {
	data, err := json.Marshal(mail)
	if err != nil {
		return err
	}

	requestAddr := a.address + "api/" + a.index + "/_doc"
	_, err = a.sendRequest(requestAddr, data)
	return err
}

// IngestBulk takes a list of mails and sends them to be ingested on bulk
//
// # The mails channel is expected to be closed by the caller
//
// chunking is how many mails will be sent, on maximum, by a single ingest request
func (a *Authentication) IngestBulk(mails chan zincindex.Mail, chunking int) error {
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

	i := 0
	requestBody := []byte("")
	for mail := range mails {
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
		i += 1

		if i%chunking == 0 {
			requestAddr := a.address + "api/_bulk"
			_, err = a.sendRequest(requestAddr, requestBody)
			if err != nil {
				return err
			}
			requestBody = []byte("")
			i = 0
		}
	}

	requestAddr := a.address + "api/_bulk"
	_, err = a.sendRequest(requestAddr, requestBody)
	return err
}

func (a *Authentication) sendRequest(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}

	req.SetBasicAuth(a.username, a.password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}

	defer resp.Body.Close()

	log.Println(resp.StatusCode)
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}

	return body, nil
}
