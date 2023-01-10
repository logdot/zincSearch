package zincSearch

import (
	"github.com/logdot/zincSearch/internal/zincindex"
	"log"
)

// IndexDB indexes the enron database at the given path and ingests it into ZincSearch.
// path is expected to point at a directory containing `maildir/`.
func (a *Authentication) IndexDB(path string) {
	log.Println("Starting indexing")

	mails := zincindex.IndexFolder(path)

	err := a.IngestBulk(mails, 10000)

	if err != nil {
		log.Println(err)
	}
}
