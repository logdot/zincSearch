package zincSearch

import (
	"encoding/json"
	"github.com/logdot/zincSearch/internal/zincindex"
	"log"
)

type SearchType = string

const (
	Matchall    SearchType = "matchall"
	Match       SearchType = "match"
	Matchprase  SearchType = "matchphrase"
	Term        SearchType = "term"
	Querystring SearchType = "querystring"
	Prefix      SearchType = "prefix"
	Wildcard    SearchType = "wildcard"
	Fuzzy       SearchType = "fuzzy"
	Daterange   SearchType = "daterange"
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

func (a *Authentication) Search(searchType SearchType, query_ SearchQuery, from uint, max_results uint, sources []string) (*SearchResult, error) {
	request := query{
		searchType,
		query_,
		from,
		max_results,
		sources,
	}

	url := a.address + "api/" + a.index + "/_search"
	body, err := json.Marshal(request)
	if err == nil {
		return nil, err
	}

	_, err = a.sendRequest(url, body)

	return nil, err
}
