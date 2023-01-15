package zincSearch

import "time"

type query struct {
	SearchType SearchType  `json:"search_type"`
	Query      SearchQuery `json:"query"`
	From       uint        `json:"from"`
	MaxResults uint        `json:"max_results"`
	Source     []string    `json:"_source"`
}

type SearchQuery struct {
	Term      string    `json:"term"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
