package models

type Release struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Status  string `json:"status"`
	Year    int    `json:"year"`
	Country string `json:"country"`

	Released          string `json:"released"`
	ReleasedFormatted string `json:"released_formatted"`

	ResourceURL string `json:"resource_url"`
	URI         string `json:"uri"`
	Thumb       string `json:"thumb"`

	DataQuality string `json:"data_quality"`
	Notes       string `json:"notes"`

	Artists      []Artist `json:"artists"`
	ArtistsSort  string   `json:"artists_sort"`
	ExtraArtists []Artist `json:"extraartists,omitempty"`

	Labels    []Label   `json:"labels"`
	Companies []Company `json:"companies"`
	Formats   []Format  `json:"formats"`

	Genres []string `json:"genres"`
	Styles []string `json:"styles"`

	Tracklist []Track `json:"tracklist"`

	Images []Image `json:"images"`
	Videos []Video `json:"videos"`

	Identifiers []Identifier `json:"identifiers"`

	Community Community `json:"community"`

	MasterID  int    `json:"master_id"`
	MasterURL string `json:"master_url"`

	NumForSale   int     `json:"num_for_sale"`
	LowestPrice  float64 `json:"lowest_price"`
	FormatQty    int     `json:"format_quantity"`
	EstimatedWgt int     `json:"estimated_weight"`

	DateAdded   string `json:"date_added"`
	DateChanged string `json:"date_changed"`

	BlockedFromSale bool `json:"blocked_from_sale"`
	IsOffensive     bool `json:"is_offensive"`
}
