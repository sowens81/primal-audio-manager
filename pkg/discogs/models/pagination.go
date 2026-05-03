package models

type Pagination struct {
	Page    int  `json:"page"`
	Pages   int  `json:"pages"`
	PerPage int  `json:"per_page"`
	Items   int  `json:"items"`
	Urls    Urls `json:"urls,omitempty"`
}

type Urls struct {
	Next string `json:"next,omitempty"`
	Last string `json:"last,omitempty"`
}
