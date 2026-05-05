package models

type Company struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	CatNo          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ResourceURL    string `json:"resource_url"`
}
