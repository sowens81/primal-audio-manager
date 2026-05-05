package models

type Image struct {
	Type        string `json:"type"`
	URI         string `json:"uri"`
	URI150      string `json:"uri150"`
	ResourceURL string `json:"resource_url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}
