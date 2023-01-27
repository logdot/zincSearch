package zincSearch

// Total contains the total amount of hits
type Total struct {
	Value int `json:"value"`
}

// Hit represents an actual hit result
type Hit struct {
	Index     string  `json:"_index"`     // the index the result is from
	Type      string  `json:"_type"`      // the type of result (seems to be identical to Index)
	Id        string  `json:"_id"`        // the id of the result
	Score     float32 `json:"_score"`     // the score of the result
	Timestamp string  `json:"@timestamp"` // a timestamp
	Source    any     `json:"_source"`    // the actual result
}

// Hits stores the hit results from a search operation
type Hits struct {
	Total Total `json:"total"` // the total amount of hits
	Hits  []Hit `json:"hits"`  // a list of Hit representing all the hits
}

// SearchResult represents the search results from a search operation in ZincSearch
type SearchResult struct {
	Took     string  `json:"took"`      // the amount of time that it took to do the search
	TimedOut bool    `json:"timed_out"` // if the search timed out or not
	MaxScore float32 `json:"max_score"` // the maximum score out of all search results
	Hits     Hits    `json:"hits"`      // a Hits struct representing the search hits
	Buckets  any     `json:"buckets"`   // the buckets that were searched
	Error    string  `json:"error"`     // any errors that elapsed. Empty if no errors
}
