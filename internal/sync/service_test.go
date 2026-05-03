package sync

import (
	"bytes"
	"errors"
	"testing"

	"github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

//
// Mock Client
//

type mockCollectionClient struct {
	GetFoldersFunc        func(string) (*models.CollectionFolders, error)
	GetFolderByIdFunc     func(string, int) (*models.CollectionFolder, error)
	GetFolderReleasesFunc func(string, int) (*models.CollectionReleases, error)
	AddFolderFunc         func(string, string) (*models.CollectionFolder, error)
}

func (m *mockCollectionClient) GetFolders(username string) (*models.CollectionFolders, error) {
	return m.GetFoldersFunc(username)
}

func (m *mockCollectionClient) GetFolderById(username string, id int) (*models.CollectionFolder, error) {
	return m.GetFolderByIdFunc(username, id)
}

func (m *mockCollectionClient) GetFolderReleases(username string, id int) (*models.CollectionReleases, error) {
	return m.GetFolderReleasesFunc(username, id)
}

func (m *mockCollectionClient) AddFolder(username string, name string) (*models.CollectionFolder, error) {
	return m.AddFolderFunc(username, name)
}

func TestService_SyncCollection(t *testing.T) {
	buf := &bytes.Buffer{}

	mock := &mockCollectionClient{
		GetFoldersFunc: func(username string) (*models.CollectionFolders, error) {
			return &models.CollectionFolders{
				Folders: []models.CollectionFolder{
					{ID: 1, Name: "Test"},
				},
			}, nil
		},
		GetFolderByIdFunc: func(username string, id int) (*models.CollectionFolder, error) {
			return &models.CollectionFolder{
				ID:    id,
				Name:  "Test",
				Count: 1,
			}, nil
		},
		GetFolderReleasesFunc: func(username string, id int) (*models.CollectionReleases, error) {
			return &models.CollectionReleases{
				Releases: []models.CollectionItem{
					{
						ID: 1,
						BasicInfo: models.BasicInformation{
							Title: "Album",
							Year:  2020,
							Artists: []models.Artist{
								{Name: "Artist"},
							},
							Genres: []string{"Rock"},
						},
					},
				},
			}, nil
		},
	}

	service := NewService(mock, "user", buf)

	err := service.SyncCollection()
	if err != nil {
		t.Fatal(err)
	}

	output := buf.String()

	if !contains(output, "Collection folders") {
		t.Fatal("expected folders output")
	}

	if !contains(output, "Album") {
		t.Fatal("expected release output")
	}
}

//
// AddFolder
//

func TestService_AddFolder(t *testing.T) {
	buf := &bytes.Buffer{}

	mock := &mockCollectionClient{
		AddFolderFunc: func(username, name string) (*models.CollectionFolder, error) {
			return &models.CollectionFolder{
				ID:    1,
				Name:  name,
				Count: 0,
			}, nil
		},
	}

	service := NewService(mock, "user", buf)

	err := service.AddFolder("Hardcore")
	if err != nil {
		t.Fatal(err)
	}

	if !contains(buf.String(), "Added folder") {
		t.Fatal("expected output")
	}
}

func TestService_AddFolder_Error(t *testing.T) {
	mock := &mockCollectionClient{
		AddFolderFunc: func(username, name string) (*models.CollectionFolder, error) {
			return nil, errors.New("fail")
		},
	}

	service := NewService(mock, "user", &bytes.Buffer{})

	err := service.AddFolder("Hardcore")
	if err == nil {
		t.Fatal("expected error")
	}
}

//
// GetFolderByID
//

func TestService_GetFolderByID(t *testing.T) {
	buf := &bytes.Buffer{}

	mock := &mockCollectionClient{
		GetFolderByIdFunc: func(username string, id int) (*models.CollectionFolder, error) {
			return &models.CollectionFolder{
				ID:    id,
				Name:  "Test",
				Count: 5,
			}, nil
		},
	}

	service := NewService(mock, "user", buf)

	err := service.GetFolderByID(1)
	if err != nil {
		t.Fatal(err)
	}

	if !contains(buf.String(), "Retrieved folder") {
		t.Fatal("expected output")
	}
}

//
// 🔹 GetFolderReleases
//

func TestService_GetFolderReleases(t *testing.T) {
	buf := &bytes.Buffer{}

	mock := &mockCollectionClient{
		GetFolderReleasesFunc: func(username string, id int) (*models.CollectionReleases, error) {
			return &models.CollectionReleases{
				Releases: []models.CollectionItem{
					{
						ID: 1,
						BasicInfo: models.BasicInformation{
							Title: "Album",
							Year:  2021,
							Artists: []models.Artist{
								{Name: "Artist"},
							},
							Genres: []string{"Electronic"},
						},
					},
				},
			}, nil
		},
	}

	service := NewService(mock, "user", buf)

	err := service.GetFolderReleases(1)
	if err != nil {
		t.Fatal(err)
	}

	output := buf.String()

	if !contains(output, "Album") {
		t.Fatal("expected release output")
	}
}

//
// Helper function to check if a substring is in a string
//

func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
