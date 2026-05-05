package models

type ExtraArtist struct {
	Name        string `json:"name"`
	Anv         string `json:"anv"`
	Join        string `json:"join"`
	Role        string `json:"role"`
	Tracks      string `json:"tracks"`
	ID          int    `json:"id"`
	ResourceURL string `json:"resource_url"`
}
