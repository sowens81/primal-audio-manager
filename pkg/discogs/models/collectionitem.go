package models

type CollectionItem struct {
	ID         int              `json:"id"`
	InstanceID int              `json:"instance_id"`
	FolderID   int              `json:"folder_id"`
	Rating     int              `json:"rating"`
	BasicInfo  BasicInformation `json:"basic_information"`
	Notes      []CollectionNote `json:"notes,omitempty"`
}
