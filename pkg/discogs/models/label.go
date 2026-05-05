package models

type Label struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	CatNo          string `json:"catno"`
	ResourceURL    string `json:"resource_url"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name,omitempty"`
	ThumbnailURL   string `json:"thumbnail_url,omitempty"`
}
