package zincSearch

import "time"

// query is an internal struct for creating the body of the search request.
type query struct {
	SearchType SearchType  `json:"search_type"`
	Query      SearchQuery `json:"query"`
	From       uint        `json:"from"`
	MaxResults uint        `json:"max_results"`
	Source     []string    `json:"_source"`
}

// SearchQuery represents a search request for ZincSearch
type SearchQuery struct {
	Term      string    `json:"term"`       // the term to search for
	StartTime time.Time `json:"start_time"` // the start time for applicable documents
	EndTime   time.Time `json:"end_time"`   // the end time for applicable documents
}
