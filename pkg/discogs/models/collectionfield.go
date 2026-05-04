package models

type CollectionField struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Position int    `json:"position"`
	Public   bool   `json:"public"`

	// Optional depending on field type
	Options []string `json:"options,omitempty"` // dropdown
	Lines   int      `json:"lines,omitempty"`   // textarea
}
