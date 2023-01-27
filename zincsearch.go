package zincSearch

import (
	"encoding/json"
	"github.com/logdot/zincSearch/internal/zincindex"
	"log"
)

type SearchType = string

// The types of search ZincSearch is capable of.
// For a complete overview see https://docs.zincsearch.com/api/search/types/
const (
	Matchall    SearchType = "matchall"    // returns all documents indiscriminately
	Match       SearchType = "match"       // is like a Term query but the input is analysed
	Matchprase  SearchType = "matchphrase" // is like a phrase query but the input is analysed
	Term        SearchType = "term"        // performs an exact match
	Querystring SearchType = "querystring" // uses the ZincSearch query language for the search
	Prefix      SearchType = "prefix"      // searches for terms that start with the provided prefix
	Wildcard    SearchType = "wildcard"    // is an alias for Prefix
	Fuzzy       SearchType = "fuzzy"       // uses fuzzy searching
	Daterange   SearchType = "daterange"   // finds documents with a date value in the specified range
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

// Search will search ZincSearch for a given query and return a SearchResult.
//
// searchType is an enum value of type SearchType that directs ZincSearch how to conduct the search.
// All the search types and what they do can be found at https://docs.zincsearch.com/api/search/types/.
//
// query_ is the actual SearchQuery to conduct.
//
// from is what result to start grabing from.
//
// max_results is how many results to return.
// This can be used in conjunction with from for pagination.
//
// sources is a list of "columns".
// If left blank it defaults to all columns.
// The ZincSearch api documentation does not specify whether this is only for return purposes, or if it also affects the searching.
func (a *Authentication) Search(searchType SearchType, query_ SearchQuery, from uint, max_results uint, sources []string) (SearchResult, error) {
	request := query{
		searchType,
		query_,
		from,
		max_results,
		sources,
	}

	url := a.address + "api/" + a.index + "/_search"
	body, err := json.Marshal(request)
	if err != nil {
		return SearchResult{}, err
	}

	response, err := a.sendRequest(url, body)

	var searchResult SearchResult
	err = json.Unmarshal(response, &searchResult)

	return searchResult, err
}
