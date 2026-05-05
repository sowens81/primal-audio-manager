package models

type Track struct {
	Position     string   `json:"position"`
	Type         string   `json:"type_"`
	Title        string   `json:"title"`
	Duration     string   `json:"duration"`
	ExtraArtists []Artist `json:"extraartists,omitempty"`
}
