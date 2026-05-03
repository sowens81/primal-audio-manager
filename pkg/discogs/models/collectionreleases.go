package models

type CollectionReleases struct {
	Pagination Pagination       `json:"pagination"`
	Releases   []CollectionItem `json:"releases"`
}
