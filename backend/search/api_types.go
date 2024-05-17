package search

// TODO: Put the query made in ClubSearchQuery.toJSON() in here

type SearchEndpointRequest struct {
	Query struct {
		QueryString struct {
			Query string `json:"query"`
		} `json:"query_string"`
	} `json:"query"`
}

type SearchEndpointResponse struct {
	Hits struct {
		Hits []struct {
			Index string  `json:"_index"`
			Score float32 `json:"_score"`
			Id    string  `json:"_id"`
		} `json:"hits"`
	} `json:"hits"`
}
