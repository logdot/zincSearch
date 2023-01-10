package zincSearch

import (
	"bytes"
	"encoding/json"
	"github.com/logdot/zincSearch/internal/zincindex"
	"io"
	"log"
	"net/http"
)

// ingestSingle takes a mail and sends it to be ingested
func (a *Authentication) IngestSingle(mail zincindex.Mail) error {
	data, err := json.Marshal(mail)
	if err != nil {
		return err
	}

	requestAddr := a.address + "api/" + a.index + "/_doc"
	return a.sendRequest(requestAddr, data)
}

func (a *Authentication) IngestBulk(mails []zincindex.Mail, chunking int) error {
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
