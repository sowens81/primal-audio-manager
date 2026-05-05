package services

import (
	"fmt"

	"github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

type ReleaseService struct {
	client APIClient
}

func NewReleaseService(client APIClient) *ReleaseService {
	return &ReleaseService{client: client}
}

//
// Discogs Release
//

// Discogs API: /releases/{release_id}
//
// The Release resource represents a particular physical or digital object released by one or more Artists.
func (s *ReleaseService) GetReleaseById(releaseID int) (models.Release, error) {
	path := fmt.Sprintf("/releases/%d", releaseID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return models.Release{}, err
	}

	var result models.Release
	err = s.client.Execute(req, &result)
	if err != nil {
		return models.Release{}, err
	}

	return result, nil
}
