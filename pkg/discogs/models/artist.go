package models

type Artist struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ResourceURL  string `json:"resource_url"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	ANV    string `json:"anv"`
	Join   string `json:"join"`
	Role   string `json:"role"`
	Tracks string `json:"tracks"`
}
