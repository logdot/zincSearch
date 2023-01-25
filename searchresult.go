package zincSearch

type Total struct {
	Value int `json:"value"`
}

type Hit struct {
	Index     string  `json:"_index"`
	Type      string  `json:"_type"`
	Id        string  `json:"_id"`
	Score     float32 `json:"_score"`
	Timestamp string  `json:"@timestamp"`
	Source    []any   `json:"_source"`
}

type Hits struct {
	Total Total `json:"total"`
	Hits  []Hit `json:"hits"`
}

type SearchResult struct {
	Took     string  `json:"took"`
	TimedOut bool    `json:"timed_out"`
	MaxScore float32 `json:"max_score"`
	Hits     Hits    `json:"hits"`
	Buckets  any     `json:"buckets"`
	Error    string  `json:"error"`
}
