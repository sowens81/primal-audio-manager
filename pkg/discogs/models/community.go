package models

type Community struct {
	Have        int    `json:"have"`
	Want        int    `json:"want"`
	Status      string `json:"status"`
	DataQuality string `json:"data_quality"`

	Rating       Rating        `json:"rating"`
	Submitter    UserSummary   `json:"submitter"`
	Contributors []UserSummary `json:"contributors"`
}
