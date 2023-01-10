package zincindex

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// IndexFolder takes a path to the enron database and indexes all of its directories and files recursively.
func IndexFolder(path string) []Mail {
	var mails []Mail

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		mail, err := IndexFile(path)
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
