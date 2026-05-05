package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

type CollectionService struct {
	client APIClient
}

func NewCollectionService(client APIClient) *CollectionService {
	return &CollectionService{client: client}
}

//
// Discogs Collection
//

// Discogs API: GET /users/{username}/collection/folders
//
// Retrieve a list of folders in a user’s collection. If the collection has been made private by its owner, authentication as the collection
// owner is required. If you are not authenticated as the collection owner, only folder ID 0 (the “All” folder) will be visible
// (if the requested user’s collection is public).
func (s *CollectionService) GetFolders(username string) (models.CollectionFoldersResponse, error) {
	path := fmt.Sprintf("/users/%s/collection/folders", username)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return models.CollectionFoldersResponse{}, err
	}

	var result models.CollectionFoldersResponse
	err = s.client.Execute(req, &result)
	if err != nil {
		return models.CollectionFoldersResponse{}, err
	}

	return result, nil
}

// Discogs API: POST /users/{username}/collection/folders
//
// Create a new folder in a user’s collection.
func (s *CollectionService) AddFolder(username string, folderName string) (models.CollectionFolder, error) {
	path := fmt.Sprintf("/users/%s/collection/folders", username)

	payload := map[string]string{
		"name": folderName,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return models.CollectionFolder{}, err
	}

	req, err := s.client.NewRequest("POST", path)
	if err != nil {
		return models.CollectionFolder{}, err
	}

	req.Body = io.NopCloser(bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	var result models.CollectionFolder
	err = s.client.Execute(req, &result)
	if err != nil {
		return models.CollectionFolder{}, err
	}

	return result, nil
}

//
// Discogs Collection Folder
//

// Discogs API: GET /users/{username}/collection/folders/{folder_id}
//
// Retrieve metadata about a folder in a user’s collection.
func (s *CollectionService) GetFolderById(username string, folderID int) (models.CollectionFolder, error) {
	path := fmt.Sprintf("/users/%s/collection/folders/%d", username, folderID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return models.CollectionFolder{}, err
	}

	var result models.CollectionFolder
	err = s.client.Execute(req, &result)
	if err != nil {
		return models.CollectionFolder{}, err
	}

	return result, nil
}

//
// Discogs Collection Items By Release
//

// Discogs API: GET /users/{username}/collection/releases/{release_id}
//
// View the user’s collection folders which contain a specified release. This will also show information about each release instance.

func (s *CollectionService) GetItemsByReleaseId(username string, releaseID int) (models.CollectionReleases, error) {
	path := fmt.Sprintf("/users/%s/collection/releases/%d", username, releaseID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return models.CollectionReleases{}, err
	}
	var result models.CollectionReleases
	err = s.client.Execute(req, &result)
	if err != nil {
		return models.CollectionReleases{}, err
	}

	return result, nil
}

//
// Discogs Collection Items By Folder
//

// Discogs API: GET /users/{username}/collection/folders/{folder_id}/releases
//
// Returns the list of item in a folder in a user’s collection. Accepts Pagination parameters.
func (s *CollectionService) GetItemsByFolder(username string, folderID int, pageOptions models.PageSettings) (models.CollectionReleases, error) {
	path := fmt.Sprintf("/users/%s/collection/folders/%d/releases?page=%d&per_page=%d", username, folderID, pageOptions.Page, pageOptions.PerPage)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return models.CollectionReleases{}, err
	}

	var result models.CollectionReleases
	err = s.client.Execute(req, &result)
	if err != nil {
		return models.CollectionReleases{}, err
	}

	return result, nil
}

//
// Discogs List Custom Fields
//

// Discogs API: GET /users/{username}/collection/fields
//
// Retrieve a list of user-defined collection notes fields. These fields are available on every release in the collection.

func (s *CollectionService) GetCustomFields(username string) (models.CollectionFields, error) {
	path := fmt.Sprintf("/users/%s/collection/fields", username)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return models.CollectionFields{}, err
	}

	var result models.CollectionFields
	err = s.client.Execute(req, &result)
	if err != nil {
		return models.CollectionFields{}, err
	}

	return result, nil
}
