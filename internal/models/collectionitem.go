package models

import "github.com/sowens81/primal-audio-manager/pkg/discogs/models"

type CollectionItem struct {
	ID                  int                     `json:"id"`
	CollectionID        int                     `json:"collection_id"`
	InstanceID          int                     `json:"instance_id"`
	ReleaseID           int                     `json:"release_id"`
	CatalogNumber       []string                `json:"catalog_number"`
	Artists             string                  `json:"artists"`
	Title               string                  `json:"title"`
	Year                int                     `json:"year"`
	Genre               []string                `json:"genre"`
	Labels              []string                `json:"labels"`
	Rating              int                     `json:"rating"`
	FolderID            int                     `json:"folder_id"`
	Notes               []models.CollectionNote `json:"notes"`
	CoverImage          string                  `json:"cover_image"`
	CoverImageThumbnail string                  `json:"cover_image_thumbnail"`
	TrackList           []models.Track          `json:"tracks"`
}
