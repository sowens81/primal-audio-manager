package models

type CollectionFolder struct {
	ID          int    `json:"id"`
	Count       int    `json:"count"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}
