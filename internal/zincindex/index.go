package zincindex

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// IndexFolder takes a path to the enron database and indexes all of its directories and files recursively.
// The returned channel is closed once all files are indexed.
//
// concurrency is the amount of files to index concurrently.
func IndexFolder(path string, concurrency uint) chan Mail {
	var wg sync.WaitGroup
	retChan := make(chan Mail)
	waitChan := make(chan any, concurrency)

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			waitChan <- struct{}{}
			mail, err := IndexFile(path)

			if err == nil {
				retChan <- mail
			}
			<-waitChan
			wg.Done()
		}()

		return nil
	})

	go func() {
		wg.Wait()
		close(retChan)
	}()

	if err != nil {
		log.Println(err)
	}

	return retChan
}

// IndexFile indexes a single file
func IndexFile(path string) (Mail, error) {
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
