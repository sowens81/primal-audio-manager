package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

type APIClient interface {
	NewRequest(method, path string) (*http.Request, error)
	Execute(req *http.Request, v interface{}) error
}

type CollectionService struct {
	client APIClient
}

func NewCollectionService(client APIClient) *CollectionService {
	return &CollectionService{client: client}
}

func (s *CollectionService) GetFolders(username string) (*models.CollectionFolders, error) {
	path := fmt.Sprintf("/users/%s/collection/folders", username)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var result models.CollectionFolders
	err = s.client.Execute(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CollectionService) GetFolderById(username string, folderID int) (*models.CollectionFolder, error) {
	path := fmt.Sprintf("/users/%s/collection/folders/%d", username, folderID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var result models.CollectionFolder
	err = s.client.Execute(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CollectionService) AddFolder(username string, folderName string) (*models.CollectionFolder, error) {
	path := fmt.Sprintf("/users/%s/collection/folders", username)

	payload := map[string]string{
		"name": folderName,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", path)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	var result models.CollectionFolder
	err = s.client.Execute(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CollectionService) GetFolderReleases(username string, folderID int, pageOptions models.PageSettings) (*models.CollectionReleases, error) {
	path := fmt.Sprintf("/users/%s/collection/folders/%d/releases?page=%d&per_page=%d", username, folderID, pageOptions.Page, pageOptions.PerPage)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var result models.CollectionReleases
	err = s.client.Execute(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
