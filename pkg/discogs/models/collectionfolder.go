package models

type CollectionFoldersResponse struct {
	Folders []CollectionFolder `json:"folders"`
}

type CollectionFolder struct {
	ID          int    `json:"id"`
	Count       int    `json:"count"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}
