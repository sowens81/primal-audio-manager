package models

type Label struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CatNo       string `json:"catno"`
	ResourceURL string `json:"resource_url"`
}
