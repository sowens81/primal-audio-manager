package models

type BasicInformation struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Year        int    `json:"year"`
	ResourceURL string `json:"resource_url"`
	Thumb       string `json:"thumb"`
	CoverImage  string `json:"cover_image"`

	Formats []Format `json:"formats"`
	Labels  []Label  `json:"labels"`
	Artists []Artist `json:"artists"`

	Genres []string `json:"genres"`
	Styles []string `json:"styles"`
}
